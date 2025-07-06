<template>
  <div class="tasks-page">
    <div class="page-header">
      <h2>任务管理</h2>
      <el-button type="primary" @click="showAddDialog = true">
        <el-icon><Plus /></el-icon>
        添加任务
      </el-button>
    </div>

    <el-card>
      <el-table :data="tasks" v-loading="loading" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="title" label="标题" />
        <el-table-column prop="description" label="描述" show-overflow-tooltip />
        <el-table-column prop="status" label="状态" width="120">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">
              {{ getStatusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="priority" label="优先级" width="100">
          <template #default="{ row }">
            <el-tag :type="getPriorityType(row.priority)">
              {{ getPriorityText(row.priority) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="due_date" label="截止日期" width="120">
          <template #default="{ row }">
            {{ row.due_date ? new Date(row.due_date).toLocaleDateString() : '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="student.name" label="分配给" />
        <el-table-column prop="creator.name" label="创建者" />
        <el-table-column label="操作" width="250">
          <template #default="{ row }">
            <el-button size="small" @click="editTask(row)">编辑</el-button>
            <el-button size="small" type="warning" @click="updateStatus(row)">更新状态</el-button>
            <el-button size="small" type="danger" @click="deleteTask(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 添加/编辑任务对话框 -->
    <el-dialog
      v-model="showAddDialog"
      :title="editingTask ? '编辑任务' : '添加任务'"
      width="600px"
    >
      <el-form
        ref="taskFormRef"
        :model="taskForm"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item label="标题" prop="title">
          <el-input v-model="taskForm.title" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="taskForm.description" type="textarea" :rows="3" />
        </el-form-item>
        <el-form-item label="优先级" prop="priority">
          <el-select v-model="taskForm.priority" style="width: 100%">
            <el-option label="低" value="low" />
            <el-option label="中" value="medium" />
            <el-option label="高" value="high" />
          </el-select>
        </el-form-item>
        <el-form-item label="截止日期" prop="due_date">
          <el-date-picker
            v-model="taskForm.due_date"
            type="datetime"
            style="width: 100%"
            placeholder="选择截止日期"
          />
        </el-form-item>
        <el-form-item label="分配给" prop="assigned_to">
          <el-select v-model="taskForm.assigned_to" style="width: 100%">
            <el-option
              v-for="student in students"
              :key="student.id"
              :label="student.name"
              :value="student.id"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showAddDialog = false">取消</el-button>
        <el-button type="primary" @click="saveTask" :loading="saving">
          保存
        </el-button>
      </template>
    </el-dialog>

    <!-- 更新状态对话框 -->
    <el-dialog v-model="showStatusDialog" title="更新任务状态" width="400px">
      <el-form>
        <el-form-item label="状态">
          <el-select v-model="statusForm.status" style="width: 100%">
            <el-option label="待处理" value="pending" />
            <el-option label="进行中" value="in_progress" />
            <el-option label="已完成" value="completed" />
            <el-option label="已取消" value="cancelled" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showStatusDialog = false">取消</el-button>
        <el-button type="primary" @click="confirmUpdateStatus" :loading="updating">
          确认
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import axios from 'axios'

export default {
  name: 'Tasks',
  setup() {
    const tasks = ref([])
    const students = ref([])
    const loading = ref(false)
    const saving = ref(false)
    const updating = ref(false)
    const showAddDialog = ref(false)
    const showStatusDialog = ref(false)
    const editingTask = ref(null)
    const currentTask = ref(null)
    
    const taskForm = reactive({
      title: '',
      description: '',
      priority: 'medium',
      due_date: null,
      assigned_to: null
    })

    const statusForm = reactive({
      status: ''
    })

    const rules = {
      title: [{ required: true, message: '请输入任务标题', trigger: 'blur' }],
      assigned_to: [{ required: true, message: '请选择分配对象', trigger: 'change' }]
    }

    const getStatusType = (status) => {
      const types = {
        pending: 'info',
        in_progress: 'warning',
        completed: 'success',
        cancelled: 'danger'
      }
      return types[status] || 'info'
    }

    const getStatusText = (status) => {
      const texts = {
        pending: '待处理',
        in_progress: '进行中',
        completed: '已完成',
        cancelled: '已取消'
      }
      return texts[status] || status
    }

    const getPriorityType = (priority) => {
      const types = {
        low: 'info',
        medium: 'warning',
        high: 'danger'
      }
      return types[priority] || 'info'
    }

    const getPriorityText = (priority) => {
      const texts = {
        low: '低',
        medium: '中',
        high: '高'
      }
      return texts[priority] || priority
    }

    const loadTasks = async () => {
      loading.value = true
      try {
        const response = await axios.get('/api/tasks')
        tasks.value = response.data.tasks
      } catch (error) {
        ElMessage.error('获取任务列表失败')
      } finally {
        loading.value = false
      }
    }

    const loadStudents = async () => {
      try {
        const response = await axios.get('/api/students')
        students.value = response.data.students
      } catch (error) {
        ElMessage.error('获取学生列表失败')
      }
    }

    const resetForm = () => {
      Object.assign(taskForm, {
        title: '',
        description: '',
        priority: 'medium',
        due_date: null,
        assigned_to: null
      })
      editingTask.value = null
    }

    const editTask = (task) => {
      editingTask.value = task
      Object.assign(taskForm, {
        title: task.title,
        description: task.description,
        priority: task.priority,
        due_date: task.due_date ? new Date(task.due_date) : null,
        assigned_to: task.assigned_to
      })
      showAddDialog.value = true
    }

    const taskFormRef = ref(null)
    
    const saveTask = async () => {
      try {
        await taskFormRef.value?.validate()
        saving.value = true

        const data = { ...taskForm }
        if (data.due_date) {
          data.due_date = data.due_date.toISOString()
        }

        if (editingTask.value) {
          await axios.put(`/api/tasks/${editingTask.value.id}`, data)
          ElMessage.success('任务更新成功')
        } else {
          await axios.post('/api/tasks', data)
          ElMessage.success('任务添加成功')
        }

        showAddDialog.value = false
        resetForm()
        loadTasks()
      } catch (error) {
        ElMessage.error(error.response?.data?.error || '操作失败')
      } finally {
        saving.value = false
      }
    }

    const updateStatus = (task) => {
      currentTask.value = task
      statusForm.status = task.status
      showStatusDialog.value = true
    }

    const confirmUpdateStatus = async () => {
      try {
        updating.value = true
        await axios.put(`/api/tasks/${currentTask.value.id}/status`, statusForm)
        ElMessage.success('状态更新成功')
        showStatusDialog.value = false
        loadTasks()
      } catch (error) {
        ElMessage.error('状态更新失败')
      } finally {
        updating.value = false
      }
    }

    const deleteTask = async (task) => {
      try {
        await ElMessageBox.confirm('确定要删除这个任务吗？', '确认删除', {
          type: 'warning'
        })
        
        await axios.delete(`/api/tasks/${task.id}`)
        ElMessage.success('任务删除成功')
        loadTasks()
      } catch (error) {
        if (error !== 'cancel') {
          ElMessage.error('删除失败')
        }
      }
    }

    onMounted(() => {
      loadTasks()
      loadStudents()
    })

    return {
      tasks,
      students,
      loading,
      saving,
      updating,
      showAddDialog,
      showStatusDialog,
      editingTask,
      currentTask,
      taskForm,
      taskFormRef,
      statusForm,
      rules,
      getStatusType,
      getStatusText,
      getPriorityType,
      getPriorityText,
      editTask,
      saveTask,
      updateStatus,
      confirmUpdateStatus,
      deleteTask
    }
  }
}
</script>

<style scoped>
.tasks-page {
  padding: 20px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.page-header h2 {
  margin: 0;
  color: #333;
}
</style> 