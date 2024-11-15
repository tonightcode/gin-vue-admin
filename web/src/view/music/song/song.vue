<template>
  <div>
    <warning-bar title="在资源权限中将此角色的资源权限清空 或者不包含创建者的角色 即可屏蔽此客户资源的显示" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="openDialog"
        >新增</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          align="left"
          label="ID"
          prop="ID"
          width="50"
        />
        <el-table-column
          align="left"
          label="名称"
          prop="name"
          width="120"
        />
        <el-table-column
          align="left"
          label="歌手"
          prop="singers"
          width="120"
        />
        <el-table-column
          align="left"
          label="类型"
          prop="type"
          width="120"
        >
          <template #default="scope">
            <span>{{ getTypeText(scope.row.type) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="播放"
          prop="url"
          width="120"
        />
        <el-table-column
          align="left"
          label="创建日期"
          width="180"
        >
          <template #default="scope">
            <span>{{ formatDate(scope.row.CreatedAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="操作"
          min-width="160"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              @click="updateCustomer(scope.row)"
            >变更</el-button>
            <el-popover
              v-model="scope.row.visible"
              placement="top"
              width="160"
            >
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button
                  type="primary"
                  link
                  @click="scope.row.visible = false"
                >取消</el-button>
                <el-button
                  type="primary"
                  @click="deleteCustomer(scope.row)"
                >确定</el-button>
              </div>
              <template #reference>
                <el-button
                  type="primary"
                  link
                  icon="delete"
                  @click="scope.row.visible = true"
                >删除</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      title="音乐"
    >
      <el-form
        ref="formRef"
        :model="form"
        label-width="80px"
        :rules="formRules"
      >
        <el-form-item
          label="类型"
          prop="type"
        >
          <el-select
            v-model="form.type"
          >
            <el-option
              v-for="t in songTypes"
              :key="t.value"
              :label="t.label"
              :value="t.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="名称"
          prop="name"
        >
          <el-input
            v-model="form.name"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="歌手">
          <el-select
            v-model="form.singerids"
            multiple
            filterable
            allow-create
            default-first-option
            placeholder="请选择"
          >
            <el-option
              v-for="item in singerData"
              :key="item.ID"
              :label="item.name"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="歌词">
          <el-input
            v-model="form.lyric"
            type="textarea"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="附件">
          <upload-song
            v-model:songCommon="songCommon"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button
            type="primary"
            @click="enterDialog"
          >确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createSong,
  updateSong,
  deleteSong,
  getSong,
  getSongList
} from '@/api/song'
import { getSingerList } from '@/api/singer'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/format'
import UploadSong from '@/components/upload/song.vue'

defineOptions({
  name: 'Song'
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const singerData = ref([])
const formRef = ref(null)

const songCommon = ref('')

const songTypes = ref([{
  value: 1,
  label: '歌曲'
}, {
  value: 2,
  label: '其他'
}])

const form = ref({
  name: '',
  url: '',
  singerids: 0,
  lyric: '',
  type: 1
})

const formRules = {
  name: [{ required: true, message: '名称不能为空', trigger: 'blur' }],
  type: [{ required: true, message: '类型不能为空', trigger: 'blur' }]
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getSongList({ page: page.value, pageSize: pageSize.value })
  const singers = await getSingerList({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
    // 遍历 tableData 并打印每个元素
    tableData.value.forEach((item) => {
      item.singerids = item.singerids.split(',')
      item.singers = item.singers.split(',').map(singer => singer.split('-')[1]).join(',')
    })
  }
  if (singers.code === 0) {
    singerData.value = singers.data.list
  }
}

getTableData()

const dialogFormVisible = ref(false)
const type = ref('')
const updateCustomer = async(row) => {
  const res = await getSong({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    res.data.song.singerids = res.data.song.singerids.split(',').map(Number)
    console.log(res.data.song)
    form.value = res.data.song
    dialogFormVisible.value = true
  }
}
const closeDialog = () => {
  dialogFormVisible.value = false
  form.value = {
    name: '',
    url: '',
    singerids: 0,
    lyric: '',
    type: 1
  }
}
const deleteCustomer = async(row) => {
  row.visible = false
  await this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    const res = deleteSong({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  }).catch(() => {
    this.$message({
      type: 'info',
      message: '已取消删除'
    })
  })
}
const enterDialog = async() => {
  formRef.value.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createSong(form.value)
        break
      case 'update':
        res = await updateSong(form.value)
        break
      default:
        res = await createSong(form.value)
        break
    }

    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '操作成功'
      })
      closeDialog()
      getTableData()
    }
  })
}
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

const getTypeText = (type) => {
  if (type === 1) {
    return '歌曲'
  } else if (type === 2) {
    return '其他'
  } else {
    return '未知'
  }
}

</script>
<style></style>
