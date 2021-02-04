<template>
  <a-card :bordered="false">
    <div class="display">
      <div class="custom-tree">
        <a-tree
          :tree-data="treeData"
          @check="onCheck"
          :checkable="true"
        >
          <template slot='title' slot-scope="item">
            <div class="title" @click="picker(item)">
              <div class="label">
                <div>
                  {{ $t(item.title) }}
                </div>
                <div style="max-width: 300px;display: flex;flex-flow: wrap">
                  <a-checkbox v-for="btn in item.operator" :key="value" style="margin-left: 0">
                    {{ authMaps[btn] }}
                  </a-checkbox>
                </div>
              </div>
            </div>
          </template>
        </a-tree>
      </div>
      <div class="form">
        <s-table
          ref="table"
          size="default"
          rowKey="id"
          :columns="columns"
          :data="loadData"
          :alert="false"
          showPagination="auto"
        >
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
      </div>
    </div>
  </a-card>
</template>

<script>
import { tree } from 'ant-design-vue'
import { queryMenuHighList,menuHighAdd,menuHighDel } from '@/api/menuHigh'
function customSlot(tree) {
  for(const value of tree) {
    value.scopedSlots = {
      title: 'title',
      other: 'other'
    }
    if(value.children.length) {
      customSlot(value.children)
    } else {
      value.operator = [ 'add','delete','get','update','export','import','update','export','import' ]
    }
  }
}
import IconSelector from '@/components/IconSelector'

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

import { STable } from '@/components'
import { queryRolesList,menusDel,roleModifyOperatorMenu } from '@/api/manage'
export default {
  name: 'TableList',
  components: {
    [tree.name]: tree,
    IconSelector,
    STable
  },

  data() {
    return {
      treeData: [],
      form: this.$form.createForm(this),
      parentId: '',
      parentTile: {},
      currentSelectedIcon: 'pause-circle',

      columns,
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
                i.alreadyActions = i.alreadyActionsObject.map(item => item.operator)
              })
            })
            return res.result
          })
      },
      selectedRowKeys: [],
      selectedRows: []
    };
  },
  created() {
    this.getData()
  },
  methods: {
    async handleDelete (menuId) {
      if(menuId <= 2) {
        return;
      }
      const data = await menusDel({ menuId })
      this.$refs.table.refresh()
    },
    getData() {
      queryMenuHighList().then(data => {
        customSlot(data)
        this.treeData = data
      })
    },
    handleIconChange (icon) {
      console.log('change Icon', icon)
      this.$message.info(<span>选中图标 <code>{icon}</code></span>)
      this.currentSelectedIcon = icon
    },
    save() {
      this.form.validateFields(async (errors, values) => {
        if(errors) { return }
        values.icon = this.currentSelectedIcon
        if(this.parentId) {
          values.parentId = this.parentId
        }
        await menuHighAdd(values)
        this.getData()
        console.log(values)
      })
    },
    picker(data) {
      this.parentId = data.id
      this.parentTile = data.title
      this.$nextTick(() => {
        this.form.setFieldsValue({ parentTitle: this.$t(this.parentTile)  })
      })

    },
    async del(id) {
      await menuHighDel({ id })
      this.getData()
    },
    primaryMenu() {
      this.parentTile = ''
      this.parentId = ''
    },
    onCheck(checkedKeys, info) {
      console.log('onCheck', checkedKeys, info);
    },
  },
}
</script>

<style lang="less" >
.display /deep/{
  display: flex;
  .form {
    flex: 1;
  }
  .custom-tree {
    flex: 0 0 400px;
    .title {
      align-items: center;
      .label {
        margin-right: 10px;
      }
    }
  }
   .ant-tree-node-content-wrapper {
     height: auto  !important;
   }
}
.search {
  margin-bottom: 54px;
}

.fold {
  width: calc(100% - 216px);
  display: inline-block
}

.operator {
  margin-bottom: 18px;
}

@media screen and (max-width: 900px) {
  .fold {
    width: 100%;
  }
}
</style>
