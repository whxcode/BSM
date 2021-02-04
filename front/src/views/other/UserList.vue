<template>
  <page-header-wrapper>
  <a-card :bordered="false">
    <div class="table-page-search-wrapper">
      <a-form layout="inline">
        <a-row :gutter="48">
          <a-col :md="8" :sm="24">
            <a-form-item label="角色ID 11">
              <a-input placeholder="请输入"/>
            </a-form-item>
          </a-col>
          <a-col :md="8" :sm="24">
            <a-form-item label="状态">
              <a-select placeholder="请选择" default-value="0">
                <a-select-option value="0">全部</a-select-option>
                <a-select-option value="1">关闭</a-select-option>
                <a-select-option value="2">运行中</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :md="8" :sm="24">
            <span class="table-page-search-submitButtons">
              <a-button type="primary" @click="$refs.table.refresh()">查询</a-button>
                 <a-button type="primary" @click="userId = '',visible = true,form.resetFields()" >添加</a-button>
              <a-button style="margin-left: 8px">重置</a-button>
            </span>
          </a-col>
        </a-row>
      </a-form>
    </div>

    <s-table
      row-key="id"
      size="default"
      ref="table"
      :columns="columns"
      :data="loadData"
    >

      <span slot="action" slot-scope="text, record">
        <a @click="handleEdit(record)">编辑</a>
        <a-divider type="vertical" />
       <a @click="handleEdit(record)">删除</a>
      </span>
    </s-table>

    <a-modal
      title="操作"
      style="top: 20px;"
      :width="800"
      v-model="visible"
      :confirm-loading="confirmLoading"
      @ok="handleOk"
    >
      <a-spin :spinning="loading">
      <a-form class="permission-form" :form="form">

        <a-form-item
          :labelCol="labelCol"
          :wrapperCol="wrapperCol"
          label="用户名称"
        >
          <a-input
            placeholder="用户名字"
            v-decorator="['userName']"
          />
        </a-form-item>
        <a-form-item
          v-if="!userId"
          :labelCol="labelCol"
          :wrapperCol="wrapperCol"
          label="账号"
        >
          <a-input
            placeholder="账号"
            v-decorator="['account']"
          />
        </a-form-item>
        <a-form-item
          v-if="!userId"
          :labelCol="labelCol"
          :wrapperCol="wrapperCol"
          label="密码"
        >
          <a-input
            placeholder="密码"
            v-decorator="['password']"
          />
        </a-form-item>

        <a-form-item label="角色名称" :labelCol="labelCol"
                     :wrapperCol="wrapperCol">
          <a-select
            v-decorator="[
            'roleId',
            { rules: [{ required: true, message: '功能权限' }] },
          ]"
            placeholder="菜单权限Id"
          >
            <a-select-option v-for="item in roles" :key="item.id">
              {{ item.label }}
            </a-select-option>
          </a-select>
        </a-form-item>


      </a-form>
      </a-spin>
    </a-modal>

  </a-card>
  </page-header-wrapper>
</template>

<script>

import { STable } from '@/components'
import { queryUsersList,getUserSingleInfo,queryRolesList,userAdd } from '@/api/manage'
const columns = [
  {
    title: '用户Id',
    dataIndex: 'id'
  },
  {
    title: '用户名称',
    dataIndex: 'userName'
  },
  {
    title: '角色Id',
    dataIndex: 'roleId'
  },
  {
    title: '角色名称',
    dataIndex: 'roleName'
  },
  {
    title: '用户账号',
    dataIndex: 'account',
  },
  {
    title: '创建时间',
    dataIndex: 'createTime',
  },
  {
    title: '更新时间',
    dataIndex: 'updateTime',
  },
  {
    title: '操作',
    width: '150px',
    dataIndex: 'action',
    scopedSlots: { customRender: 'action' }
  }
]

export default {
  name: 'TableList',
  components: {
    STable
  },
  data () {
    return {
      visible: false,
      loading: false,
      confirmLoading: false,
      labelCol: {
        xs: { span: 24 },
        sm: { span: 5 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 16 }
      },
      roles: [],
      userId: null,
      form: this.$form.createForm(this),
      permissions: [],
      // 查询参数
      queryParam: {},
      // 表头
      columns,
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        return queryUsersList(parameter)
          .then(res => {
            // 展开全部行
            this.expandedRowKeys = res.result.data.map(item => item.id)
            return res.result
          })
      },

      expandedRowKeys: [],
      selectedRowKeys: [],
      selectedRows: []
    }
  },
  created() {
    queryRolesList({ pageNo: 1,pageSize: 10000 })
    .then(({ result: { data } }) => {
      this.roles = data
      console.log(data)
    })
  },
  methods: {
    async handleEdit (record) { // 获取用户详情
      this.visible = true
      this.loading = true
      const data = await getUserSingleInfo({ id: record.id })
      this.loading = false
      this.userId = data.id
      this.form.setFieldsValue(data)

    },
    handleOk (e) {
      e.preventDefault()
      this.form.validateFields(async (err, values) => {
        if(this.userId) {
          values.id = this.userId
        }
        this.confirmLoading = true
        await userAdd(values)
        this.confirmLoading = false
        this.visible = false
        this.$refs.table.refresh()
      })
    },

  },
}
</script>

<style lang="less" scoped>
.permission-form {
  /deep/ .permission-group {
    margin-top: 0;
    margin-bottom: 0;
  }
}

</style>
