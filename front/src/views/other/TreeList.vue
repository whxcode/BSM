<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <div class="table-page-search-wrapper">
        <a-form layout="inline">
          <a-row :gutter="48">
            <a-col :md="8" :sm="24">
              <a-form-item label="规则编号">
                <a-input v-model="queryParam.id" placeholder=""/>
              </a-form-item>
            </a-col>
            <a-col :md="8" :sm="24">
              <a-form-item label="使用状态">
                <a-select v-model="queryParam.status" placeholder="请选择" default-value="0">
                  <a-select-option value="0">全部</a-select-option>
                  <a-select-option value="1">关闭</a-select-option>
                  <a-select-option value="2">运行中</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <template v-if="advanced">
              <a-col :md="8" :sm="24">
                <a-form-item label="调用次数">
                  <a-input-number v-model="queryParam.callNo" style="width: 100%"/>
                </a-form-item>
              </a-col>
              <a-col :md="8" :sm="24">
                <a-form-item label="更新日期">
                  <a-date-picker v-model="queryParam.date" style="width: 100%" placeholder="请输入更新日期"/>
                </a-form-item>
              </a-col>
              <a-col :md="8" :sm="24">
                <a-form-item label="使用状态">
                  <a-select v-model="queryParam.useStatus" placeholder="请选择" default-value="0">
                    <a-select-option value="0">全部</a-select-option>
                    <a-select-option value="1">关闭</a-select-option>
                    <a-select-option value="2">运行中</a-select-option>
                  </a-select>
                </a-form-item>
              </a-col>
              <a-col :md="8" :sm="24">
                <a-form-item label="使用状态">
                  <a-select placeholder="请选择" default-value="0">
                    <a-select-option value="0">全部</a-select-option>
                    <a-select-option value="1">关闭</a-select-option>
                    <a-select-option value="2">运行中</a-select-option>
                  </a-select>
                </a-form-item>
              </a-col>
            </template>
            <a-col :md="!advanced && 8 || 24" :sm="24">
              <span class="table-page-search-submitButtons" :style="advanced && { float: 'right', overflow: 'hidden' } || {} ">
                <a-button type="primary" @click="$refs.table.refresh(true)">查询</a-button>
                <a-button style="margin-left: 8px" @click="() => this.queryParam = {}">重置</a-button>
                <a @click="toggleAdvanced" style="margin-left: 8px">
                  {{ advanced ? '收起' : '展开' }}
                  <a-icon :type="advanced ? 'up' : 'down'"/>
                </a>
              </span>
            </a-col>
          </a-row>
        </a-form>
      </div>
      <div class="table-operator">
        <a-button type="primary" icon="plus" @click="handleAdd">新建</a-button>
        <a-dropdown v-action:update v-if="selectedRowKeys.length > 0">
          <a-menu slot="overlay">
            <a-menu-item key="1"><a-icon type="delete" />删除</a-menu-item>
            <!-- lock | unlock -->
            <a-menu-item key="2"><a-icon type="lock" />锁定</a-menu-item>
          </a-menu>
          <a-button style="margin-left: 8px">
            批量操作 <a-icon type="down" />
          </a-button>
        </a-dropdown>
      </div>
      <s-table
        ref="table"
        size="default"
        rowKey="id"
        :columns="columns"
        :data="loadData"
        :alert="true"
        :rowSelection="rowSelection"
        showPagination="auto"
      >
        <div slot="expandedRowRender" slot-scope="record" style="margin: 0">
          <a-list :grid="{ gutter: 16, xs: 1, sm: 2, md: 4, lg: 4, xl: 6, xxl: 3 }" :data-source="record.operators">
            <a-list-item slot="renderItem" slot-scope="item, index">
              <a-card :title="item.menuName">
               <a-checkbox-group @change="(data) => {
                 onChange(data,item)
               }" v-model="item.alreadyActions">
                  <a-row>
                    <a-col :span="8" v-for="btn in JSON.parse(item.actions)" :key="btn">
                      <a-checkbox :value="btn">
                        {{ authMaps[btn] }}
                      </a-checkbox>
                    </a-col>
                  </a-row>
                </a-checkbox-group>
              </a-card>
            </a-list-item>
          </a-list>
        </div>
        <span slot="actions" slot-scope="text, record">
          <template>
            {{ text }}
          </template>
        </span>
        <span slot="action" slot-scope="text, record">
          <template>
             <a-popconfirm placement="topLeft" ok-text="Yes" cancel-text="No" @confirm="handleDelete(record.id)">
                <template slot="title">
                  <p>确定要删除吗?</p>
                </template>
                   <a>删除</a>
              </a-popconfirm>
            <a-divider type="vertical" />
            <a @click="handleSub(record)">编辑</a>
          </template>
        </span>
      </s-table>
      <create-form
        ref="createModal"
        :visible="visible"
        :loading="confirmLoading"
        :model="mdl"
        @cancel="handleCancel"
        @ok="() => { $refs.table.refresh() }"
      />
    </a-card>
  </page-header-wrapper>
</template>

<script>
import moment from 'moment'
import { STable, Ellipsis } from '@/components'
import { queryRolesList,menusDel,roleModifyOperatorMenu } from '@/api/manage'

const columns = [
  {
    title: '编号',
    dataIndex: 'id'
  },
  {
    title: '角色名称',
    dataIndex: 'label'
  },
  {
    title: '操作权限',
    dataIndex: 'label'
  },
  {
    title: '操作',
    dataIndex: 'action',
    width: '150px',
    scopedSlots: { customRender: 'action' }
  }
]
import { authMaps } from '@/config/defaultSettings'

export default {
  name: 'TableList',
  components: {
    STable,
    Ellipsis
  },
  data () {
    this.columns = columns
    return {
      authMaps,
      // create model
      visible: false,
      confirmLoading: false,
      mdl: null,
      // 高级搜索 展开/关闭
      advanced: false,
      // 查询参数
      queryParam: {},
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        const requestParameters = Object.assign({}, parameter, this.queryParam)
        // console.log('loadData request parameters:', requestParameters)
        return queryRolesList(requestParameters)
          .then(res => {
            res.result.data.forEach(item => {
              item.operators.forEach(i => {
                i.alreadyActions = JSON.parse(i.alreadyActions)
              })
            })
            return res.result
          })
      },
      selectedRowKeys: [],
      selectedRows: []
    }
  },
  filters: {
    toJS(data) { return JSON.parse(data) }
  },
  computed: {
    rowSelection () {
      return {
        selectedRowKeys: this.selectedRowKeys,
        onChange: this.onSelectChange
      }
    }
  },
  methods: {
    handleAdd () {
      this.$refs.createModal.visible = true
    },

    async handleDelete (menuId) {
      if(menuId <= 2) {
        return;
      }
      const data = await menusDel({ menuId })
      this.$refs.table.refresh()
    },
    onChange(data,item) {
      roleModifyOperatorMenu({
        operatorId: item.operatorId,
        actions: JSON.stringify(data)
      })
    },
    handleCancel () {
      this.visible = false
      const form = this.$refs.createModal.form
      form.resetFields() // 清理表单数据（可不做）
    },
    handleSub (record) {
      this.$refs.createModal.init(record.id)
    },
    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    resetSearchForm () {
      this.queryParam = {
        date: moment(new Date())
      }
    }
  }
}
</script>
