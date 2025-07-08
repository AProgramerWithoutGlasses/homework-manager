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
        <el-table-column prop="name" label="任务名称" />
        <el-table-column prop="start_time" label="开始时间" width="150">
          <template #default="{ row }">
            {{ row.start_time ? new Date(row.start_time).toLocaleString() : '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="end_time" label="结束时间" width="150">
          <template #default="{ row }">
            {{ row.end_time ? new Date(row.end_time).toLocaleString() : '-' }}
          </template>
        </el-table-column>
        <el-table-column label="进度" width="120">
          <template #default="{ row }">
            <el-progress 
              :percentage="row.students && row.students.length > 0 ? Math.round((row.completed_count / row.students.length) * 100) : 0"
              :status="getProgressStatus(row)"
            />
            <span>{{ row.completed_count }}/{{ row.students ? row.students.length : 0 }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="120">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">
              {{ getStatusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="参与学生" width="200">
          <template #default="{ row }">
            <el-tag v-for="student in row.students" :key="student.id" size="small" style="margin-right: 5px;">
              {{ student.name }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="teacher.name" label="创建老师" width="120" />
        <el-table-column label="操作" width="300">
          <template #default="{ row }">
            <el-button size="small" @click="editTask(row)">编辑</el-button>
            <el-button size="small" type="warning" @click="updateStatus(row)">更新状态</el-button>
            <el-button size="small" type="info" @click="viewTaskStudents(row)">查看学生</el-button>
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
        label-width="100px"
      >
        <el-form-item label="任务名称" prop="name">
          <el-input v-model="taskForm.name" />
        </el-form-item>
        <el-form-item label="任务描述" prop="description">
          <el-input v-model="taskForm.description" type="textarea" rows="2" placeholder="请输入任务描述（可选）" />
        </el-form-item>
        <el-form-item label="开始时间" prop="start_time">
          <el-date-picker
            v-model="taskForm.start_time"
            type="datetime"
            style="width: 100%"
            placeholder="选择开始时间"
          />
        </el-form-item>
        <el-form-item label="结束时间" prop="end_time">
          <el-date-picker
            v-model="taskForm.end_time"
            type="datetime"
            style="width: 100%"
            placeholder="选择结束时间"
          />
        </el-form-item>
        <el-form-item label="参与学生" prop="student_ids">
          <div class="student-select-toolbar">
            <el-select ref="classSelectRef" placeholder="按班级选择" :filterable="false" @click="handleSelectTrigger(classSelectRef)" style="width: 120px">
              <el-option v-for="cls in classOptions" :key="cls" :label="cls" :value="cls" @click="onClassOptionClick(cls)" />
            </el-select>

            <el-select ref="gradeSelectRef" placeholder="按年级选择" :filterable="false" @click="handleSelectTrigger(gradeSelectRef)" style="width: 120px">
              <el-option v-for="grd in gradeOptions" :key="grd" :label="grd" :value="grd" @click="onGradeOptionClick(grd)" />
            </el-select>

            <el-button @click="selectAllStudents" size="small">全选所有学生</el-button>
            <el-button @click="clearSelectedStudents" size="small">清空选择</el-button>
          </div>
          <el-select v-model="taskForm.student_ids" multiple style="width: 100%" placeholder="选择参与学生">
            <el-option
              v-for="student in students"
              :key="student.id"
              :label="`${student.name}（${student.class}，${student.grade}）`"
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
        <el-button @click="confirmUpdateStatus" :loading="updating">
          确认
        </el-button>
      </template>
    </el-dialog>

    <!-- 查看任务学生完成情况对话框 -->
    <el-dialog v-model="showStudentsDialog" title="任务学生完成情况" width="800px">
      <div v-loading="studentsLoading">
        <div class="task-info">
          <h3>{{ currentTask?.name }}</h3>
          <p>完成进度: {{ completionStats.completed_count }}/{{ completionStats.total_count }}</p>
        </div>
        
        <el-table :data="taskStudents" style="width: 100%">
          <el-table-column prop="student_id" label="学生ID" width="80" />
          <el-table-column label="学生姓名" width="120">
            <template #default="{ row }">
              {{ getStudentName(row.student_id) }}
            </template>
          </el-table-column>
          <el-table-column prop="completed" label="完成状态" width="120">
            <template #default="{ row }">
              <el-tag :type="row.completed ? 'success' : 'info'">
                {{ row.completed ? '已完成' : '未完成' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="completed_at" label="完成时间" width="180">
            <template #default="{ row }">
              {{ row.completed_at ? new Date(row.completed_at).toLocaleString() : '-' }}
            </template>
          </el-table-column>
          <el-table-column label="操作" width="150">
            <template #default="{ row }">
              <el-button 
                size="small" 
                :type="row.completed ? 'warning' : 'success'"
                @click="toggleStudentStatus(row)"
                :loading="row.updating"
              >
                {{ row.completed ? '标记未完成' : '标记完成' }}
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-dialog>

    <!-- 在el-dialog外部添加锚点和popover（班级/年级各一个） -->
    <span ref="classPopoverAnchor" style="display:inline-block;width:0;height:0;position:fixed;left:0;top:0;z-index:9999;"></span>
    <el-popover
      placement="right-start"
      width="350"
      v-model:visible="classPopoverVisible"
      :virtual-ref="classPopoverAnchor"
      virtual-triggering
      trigger="manual"
    >
      <div>
        <el-checkbox-group v-model="classPopoverSelected">
          <el-checkbox
            v-for="stu in classPopoverStudents"
            :key="stu.id"
            :label="stu.id"
          >{{ stu.name }}（{{ stu.username }}）</el-checkbox>
        </el-checkbox-group>
        <div style="margin-top: 10px;">
          <el-button @click="selectAllClassPopover" size="small">全选</el-button>
          <el-button @click="clearAllClassPopover" size="small">全不选</el-button>
          <el-button type="primary" @click="confirmClassPopover" size="small">确定</el-button>
          <el-button @click="cancelClassPopover" size="small">取消</el-button>
        </div>
      </div>
    </el-popover>
    <span ref="gradePopoverAnchor" style="display:inline-block;width:0;height:0;position:fixed;left:0;top:0;z-index:9999;"></span>
    <el-popover
      placement="right-start"
      width="350"
      v-model:visible="gradePopoverVisible"
      :virtual-ref="gradePopoverAnchor"
      virtual-triggering
      trigger="manual"
    >
      <div>
        <el-checkbox-group v-model="gradePopoverSelected">
          <el-checkbox
            v-for="stu in gradePopoverStudents"
            :key="stu.id"
            :label="stu.id"
          >{{ stu.name }}（{{ stu.username }}）</el-checkbox>
        </el-checkbox-group>
        <div style="margin-top: 10px;">
          <el-button @click="selectAllGradePopover" size="small">全选</el-button>
          <el-button @click="clearAllGradePopover" size="small">全不选</el-button>
          <el-button type="primary" @click="confirmGradePopover" size="small">确定</el-button>
          <el-button @click="cancelGradePopover" size="small">取消</el-button>
        </div>
      </div>
    </el-popover>
  </div>
</template>

<script>
import { ref, reactive, onMounted, computed, nextTick, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import axios from 'axios'

export default {
  name: 'Tasks',
  setup() {
    const tasks = ref([])
    const students = ref([])
    const taskStudents = ref([])
    const completionStats = ref({ completed_count: 0, total_count: 0 })
    const loading = ref(false)
    const saving = ref(false)
    const updating = ref(false)
    const studentsLoading = ref(false)
    const showAddDialog = ref(false)
    const showStatusDialog = ref(false)
    const showStudentsDialog = ref(false)
    const editingTask = ref(null)
    const currentTask = ref(null)
    
    const taskForm = reactive({
      name: '',
      description: '',
      start_time: null,
      end_time: null,
      student_ids: []
    })

    const statusForm = reactive({
      status: ''
    })

    const rules = {
      name: [{ required: true, message: '请输入任务名称', trigger: 'blur' }]
    }

    const selectedClass = ref('')
    const selectedGrade = ref('')
    const classOptions = computed(() => [...new Set(students.value.map(s => s.class).filter(Boolean))])
    const gradeOptions = computed(() => [...new Set(students.value.map(s => s.grade).filter(Boolean))])

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

    const getProgressStatus = (task) => {
      if (task.status === 'completed') return 'success'
      if (task.status === 'cancelled') return 'exception'
      if (task.students && task.completed_count >= task.students.length) return 'success'
      return ''
    }

    const getStudentName = (studentId) => {
      const student = students.value.find(s => s.id === studentId)
      return student ? student.name : `学生${studentId}`
    }

    const loadTasks = async () => {
      loading.value = true
      try {
        const response = await axios.get('/api/tasks')
        tasks.value = (response.data.data.tasks || []).map(t => ({
          ...t,
          id: t.ID
        }))
      } catch (error) {
        ElMessage.error('获取任务列表失败')
      } finally {
        loading.value = false
      }
    }

    const loadStudents = async () => {
      try {
        const response = await axios.get('/api/students')
        students.value = (response.data.data.students || []).map(s => ({
          ...s,
          id: s.ID
        }))
      } catch (error) {
        ElMessage.error('获取学生列表失败')
      }
    }

    const loadTaskStudents = async (taskId) => {
      studentsLoading.value = true
      try {
        const response = await axios.get(`/api/tasks/${taskId}/students`)
        taskStudents.value = response.data.task_students.map(ts => ({
          ...ts,
          updating: false
        }))
        
        // 获取完成统计
        const statsResponse = await axios.get(`/api/tasks/${taskId}/completion-stats`)
        completionStats.value = statsResponse.data
      } catch (error) {
        ElMessage.error('获取任务学生信息失败')
      } finally {
        studentsLoading.value = false
      }
    }

    const resetForm = () => {
      Object.assign(taskForm, {
        name: '',
        description: '',
        start_time: null,
        end_time: null,
        student_ids: []
      })
      editingTask.value = null
    }

    const editTask = (task) => {
      editingTask.value = task
      Object.assign(taskForm, {
        name: task.name,
        description: task.description,
        start_time: task.start_time ? new Date(task.start_time) : null,
        end_time: task.end_time ? new Date(task.end_time) : null,
        student_ids: task.students ? task.students.map(s => s.id) : []
      })
      showAddDialog.value = true
    }

    const taskFormRef = ref(null)
    
    const saveTask = async () => {
      try {
        await taskFormRef.value?.validate()
        saving.value = true

        const data = { ...taskForm }
        if (data.start_time) {
          data.start_time = data.start_time.toISOString()
        }
        if (data.end_time) {
          data.end_time = data.end_time.toISOString()
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

    const viewTaskStudents = async (task) => {
      currentTask.value = task
      showStudentsDialog.value = true
      await loadTaskStudents(task.id)
    }

    const toggleStudentStatus = async (taskStudent) => {
      taskStudent.updating = true
      try {
        await axios.put(`/api/tasks/${currentTask.value.id}/students/${taskStudent.student_id}/status`, {
          completed: !taskStudent.completed
        })
        
        taskStudent.completed = !taskStudent.completed
        if (taskStudent.completed) {
          taskStudent.completed_at = new Date().toISOString()
        } else {
          taskStudent.completed_at = null
        }
        
        // 重新加载完成统计
        const statsResponse = await axios.get(`/api/tasks/${currentTask.value.id}/completion-stats`)
        completionStats.value = statsResponse.data
        
        ElMessage.success('状态更新成功')
      } catch (error) {
        ElMessage.error('状态更新失败')
      } finally {
        taskStudent.updating = false
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

    const classSelectRef = ref(null)
    const classPopoverAnchor = ref(null)
    const classPopoverVisible = ref(false)
    const classPopoverStudents = ref([])
    const classPopoverSelected = ref([])
    
    const selectAllClassPopover = () => {
      classPopoverSelected.value = classPopoverStudents.value.map(s => s.id)
    }
    const clearAllClassPopover = () => {
      classPopoverSelected.value = []
    }
    const confirmClassPopover = () => {
      const currentTaskStudentIds = new Set(taskForm.student_ids);
      classPopoverStudents.value.forEach(s => currentTaskStudentIds.delete(s.id)); // Remove all students from current class
      classPopoverSelected.value.forEach(id => currentTaskStudentIds.add(id)); // Add newly selected students

      taskForm.student_ids = Array.from(currentTaskStudentIds);

      classPopoverVisible.value = false
      console.log('Class Popover: Set to hidden by confirm', classPopoverVisible.value)
      // 确保el-select下拉菜单关闭
      if (classSelectRef.value) {
        classSelectRef.value.blur();
      }
      selectedClass.value = ''; // 清空选择，允许下次重新触发watch
    }
    const cancelClassPopover = () => {
      classPopoverVisible.value = false
      console.log('Class Popover: Set to hidden by cancel', classPopoverVisible.value)
      // 确保el-select下拉菜单关闭
      if (classSelectRef.value) {
        classSelectRef.value.blur();
      }
      selectedClass.value = ''; // 清空选择
    }

    const gradeSelectRef = ref(null)
    const gradePopoverAnchor = ref(null)
    const gradePopoverVisible = ref(false)
    const gradePopoverStudents = ref([])
    const gradePopoverSelected = ref([])
    
    const selectAllGradePopover = () => {
      gradePopoverSelected.value = gradePopoverStudents.value.map(s => s.id)
    }
    const clearAllGradePopover = () => {
      gradePopoverSelected.value = []
    }
    const confirmGradePopover = () => {
      const currentTaskStudentIds = new Set(taskForm.student_ids);
      gradePopoverStudents.value.forEach(s => currentTaskStudentIds.delete(s.id)); // Remove all students from current grade
      gradePopoverSelected.value.forEach(id => currentTaskStudentIds.add(id)); // Add newly selected students

      taskForm.student_ids = Array.from(currentTaskStudentIds);

      gradePopoverVisible.value = false
      console.log('Grade Popover: Set to hidden by confirm', gradePopoverVisible.value)
      // 确保el-select下拉菜单关闭
      if (gradeSelectRef.value) {
        gradeSelectRef.value.blur();
      }
      selectedGrade.value = ''; // 清空选择
    }
    const cancelGradePopover = () => {
      gradePopoverVisible.value = false
      console.log('Grade Popover: Set to hidden by cancel', gradePopoverVisible.value)
      // 确保el-select下拉菜单关闭
      if (gradeSelectRef.value) {
        gradeSelectRef.value.blur();
      }
      selectedGrade.value = ''; // 清空选择
    }
    const onGradeSelectChange = (val) => {
      selectedGrade.value = val;
      gradePopoverStudents.value = students.value.filter(s => s.grade === val);
      gradePopoverSelected.value = gradePopoverStudents.value
        .filter(s => taskForm.student_ids.includes(s.id))
        .map(s => s.id);

      gradePopoverVisible.value = true;
      console.log('Grade Popover: Set to visible via @change', gradePopoverVisible.value);

      nextTick(() => {
        const selectEl = gradeSelectRef.value?.$el;
        if (selectEl) {
          const rect = selectEl.getBoundingClientRect();
          const anchor = gradePopoverAnchor.value;
          if (anchor) {
            anchor.style.left = rect.right + 'px';
            anchor.style.top = rect.top + 'px';
          }
        }
        // Re-open el-select dropdown after a short delay
        setTimeout(() => {
          if (gradeSelectRef.value) {
            gradeSelectRef.value.toggleMenu();
            console.log('Grade Select: toggleMenu called after popover appears with delay');
          }
        }, 50); // Small delay to allow el-select to complete its closing
      });
    }

    const selectAllStudents = () => {
      taskForm.student_ids = students.value.map(s => s.id)
    }

    const clearSelectedStudents = () => {
      taskForm.student_ids = []
    }

    const onClassOptionClick = (cls) => {
      selectedClass.value = cls;
      classPopoverStudents.value = students.value.filter(s => s.class === cls);
      classPopoverSelected.value = classPopoverStudents.value
        .filter(s => taskForm.student_ids.includes(s.id))
        .map(s => s.id);

      classPopoverVisible.value = true;
      console.log('Class Popover: Set to visible via @click.stop.prevent', classPopoverVisible.value);

      nextTick(() => {
        const selectEl = classSelectRef.value?.$el;
        if (selectEl) {
          const rect = selectEl.getBoundingClientRect();
          const anchor = classPopoverAnchor.value;
          if (anchor) {
            anchor.style.left = rect.right + 'px';
            anchor.style.top = rect.top + 'px';
          }
        }
        // el-select should remain open due to @click.stop.prevent on el-option
        // No explicit toggleMenu() here, rely on @click.stop.prevent
      });
    }

    const onGradeOptionClick = (grd) => {
      selectedGrade.value = grd;
      gradePopoverStudents.value = students.value.filter(s => s.grade === grd);
      gradePopoverSelected.value = gradePopoverStudents.value
        .filter(s => taskForm.student_ids.includes(s.id))
        .map(s => s.id);

      gradePopoverVisible.value = true;
      console.log('Grade Popover: Set to visible via @click.stop.prevent', gradePopoverVisible.value);

      nextTick(() => {
        const selectEl = gradeSelectRef.value?.$el;
        if (selectEl) {
          const rect = selectEl.getBoundingClientRect();
          const anchor = gradePopoverAnchor.value;
          if (anchor) {
            anchor.style.left = rect.right + 'px';
            anchor.style.top = rect.top + 'px';
          }
        }
        // el-select should remain open due to @click.stop.prevent on el-option
        // No explicit toggleMenu() here, rely on @click.stop.prevent
      });
    }

    const onClassDropdownVisibleChange = (visible) => {
      // If el-select attempts to close (!visible) and Popover is open, prevent el-select from closing
      if (!visible && classPopoverVisible.value) {
        console.log('Class Select visible-change: !visible detected, classPopoverVisible.value =', classPopoverVisible.value, '. Forcing el-select open.');
        if (classSelectRef.value) {
          classSelectRef.value.toggleMenu(); // Force re-open the dropdown
        }
      } else {
        console.log('Class Select dropdown visible change:', visible, 'Popover visible:', classPopoverVisible.value);
      }
    }

    const onGradeDropdownVisibleChange = (visible) => {
      // If el-select attempts to close (!visible) and Popover is open, prevent el-select from closing
      if (!visible && gradePopoverVisible.value) {
        console.log('Grade Select visible-change: !visible detected, gradePopoverVisible.value =', gradePopoverVisible.value, '. Forcing el-select open.');
        if (gradeSelectRef.value) {
          gradeSelectRef.value.toggleMenu(); // Force re-open the dropdown
        }
      } else {
        console.log('Grade Select dropdown visible change:', visible, 'Popover visible:', gradePopoverVisible.value);
      }
    }

    const handleSelectTrigger = (selectRef) => {
      // This method is called when the select dropdown is triggered
      // It can be used to open the dropdown programmatically
      if (selectRef) {
        selectRef.toggleMenu();
      }
    }

    onMounted(() => {
      loadTasks()
      loadStudents()
    })

    return {
      tasks,
      students,
      taskStudents,
      completionStats,
      loading,
      saving,
      updating,
      studentsLoading,
      showAddDialog,
      showStatusDialog,
      showStudentsDialog,
      editingTask,
      currentTask,
      taskForm,
      taskFormRef,
      statusForm,
      rules,
      getStatusType,
      getStatusText,
      getProgressStatus,
      getStudentName,
      editTask,
      saveTask,
      updateStatus,
      confirmUpdateStatus,
      viewTaskStudents,
      toggleStudentStatus,
      deleteTask,
      selectedClass,
      selectedGrade,
      classOptions,
      gradeOptions,
      classSelectRef,
      classPopoverAnchor,
      classPopoverVisible,
      classPopoverStudents,
      classPopoverSelected,
      selectAllClassPopover,
      clearAllClassPopover,
      confirmClassPopover,
      cancelClassPopover,
      gradeSelectRef,
      gradePopoverAnchor,
      gradePopoverVisible,
      gradePopoverStudents,
      gradePopoverSelected,
      selectAllGradePopover,
      clearAllGradePopover,
      confirmGradePopover,
      cancelGradePopover,
      selectAllStudents,
      clearSelectedStudents,
      onGradeSelectChange,
      onClassDropdownVisibleChange,
      onGradeDropdownVisibleChange,
      handleSelectTrigger
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

.task-info {
  margin-bottom: 20px;
  padding: 15px;
  background: #f5f5f5;
  border-radius: 4px;
}

.task-info h3 {
  margin: 0 0 10px 0;
  color: #333;
}

.task-info p {
  margin: 0;
  color: #666;
}

.student-select-toolbar {
  display: flex;
  gap: 8px;
  margin-bottom: 8px;
}
</style> 