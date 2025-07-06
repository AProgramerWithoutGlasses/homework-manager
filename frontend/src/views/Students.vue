<template>
  <div class="students-page">
    <div class="page-header">
      <h2>学生管理</h2>
      <el-button type="primary" @click="showAddDialog = true">
        <el-icon><Plus /></el-icon>
        添加学生
      </el-button>
    </div>

    <el-card>
      <el-table :data="students" v-loading="loading" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="姓名" />
        <el-table-column prop="email" label="邮箱" />
        <el-table-column prop="phone" label="电话" />
        <el-table-column prop="class" label="班级" />
        <el-table-column prop="grade" label="年级" />
        <el-table-column prop="status" label="状态">
          <template #default="{ row }">
            <el-tag :type="row.status === 'active' ? 'success' : 'danger'">
              {{ row.status === 'active' ? '活跃' : '非活跃' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button size="small" @click="editStudent(row)">编辑</el-button>
            <el-button size="small" type="danger" @click="deleteStudent(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 添加/编辑学生对话框 -->
    <el-dialog
      v-model="showAddDialog"
      :title="editingStudent ? '编辑学生' : '添加学生'"
      width="500px"
    >
      <el-form
        ref="studentFormRef"
        :model="studentForm"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item label="姓名" prop="name">
          <el-input v-model="studentForm.name" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="studentForm.email" />
        </el-form-item>
        <el-form-item label="电话" prop="phone">
          <el-input v-model="studentForm.phone" />
        </el-form-item>
        <el-form-item label="班级" prop="class">
          <el-input v-model="studentForm.class" />
        </el-form-item>
        <el-form-item label="年级" prop="grade">
          <el-input v-model="studentForm.grade" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="studentForm.status" style="width: 100%">
            <el-option label="活跃" value="active" />
            <el-option label="非活跃" value="inactive" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showAddDialog = false">取消</el-button>
        <el-button type="primary" @click="saveStudent" :loading="saving">
          保存
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
  name: 'Students',
  setup() {
    const students = ref([])
    const loading = ref(false)
    const saving = ref(false)
    const showAddDialog = ref(false)
    const editingStudent = ref(null)
    const studentForm = reactive({
      name: '',
      email: '',
      phone: '',
      class: '',
      grade: '',
      status: 'active'
    })

    const rules = {
      name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
      email: [
        { required: true, message: '请输入邮箱', trigger: 'blur' },
        { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
      ]
    }

    const loadStudents = async () => {
      loading.value = true
      try {
        const response = await axios.get('/api/students')
        students.value = response.data.students
      } catch (error) {
        ElMessage.error('获取学生列表失败')
      } finally {
        loading.value = false
      }
    }

    const resetForm = () => {
      Object.assign(studentForm, {
        name: '',
        email: '',
        phone: '',
        class: '',
        grade: '',
        status: 'active'
      })
      editingStudent.value = null
    }

    const editStudent = (student) => {
      editingStudent.value = student
      Object.assign(studentForm, student)
      showAddDialog.value = true
    }

    const studentFormRef = ref(null)
    
    const saveStudent = async () => {
      try {
        await studentFormRef.value?.validate()
        saving.value = true

        if (editingStudent.value) {
          await axios.put(`/api/students/${editingStudent.value.id}`, studentForm)
          ElMessage.success('学生信息更新成功')
        } else {
          await axios.post('/api/students', studentForm)
          ElMessage.success('学生添加成功')
        }

        showAddDialog.value = false
        resetForm()
        loadStudents()
      } catch (error) {
        ElMessage.error(error.response?.data?.error || '操作失败')
      } finally {
        saving.value = false
      }
    }

    const deleteStudent = async (student) => {
      try {
        await ElMessageBox.confirm('确定要删除这个学生吗？', '确认删除', {
          type: 'warning'
        })
        
        await axios.delete(`/api/students/${student.id}`)
        ElMessage.success('学生删除成功')
        loadStudents()
      } catch (error) {
        if (error !== 'cancel') {
          ElMessage.error('删除失败')
        }
      }
    }

    onMounted(() => {
      loadStudents()
    })

    return {
      students,
      loading,
      saving,
      showAddDialog,
      editingStudent,
      studentForm,
      studentFormRef,
      rules,
      editStudent,
      saveStudent,
      deleteStudent
    }
  }
}
</script>

<style scoped>
.students-page {
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