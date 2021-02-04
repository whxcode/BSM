<template>
  <a-card :bordered="false">
    <div class="display">
      <div class="custom-tree">
        <a-tree
          :tree-data="treeData"
          @check="onCheck"
        >
          <template slot='title' slot-scope="item">
            <div class="title" @click="picker(item)">
              <div class="label">
                {{ $t(item.title) }}
              </div>
              <div class="other">
                <a-tag color="pink" @click.stop>修改</a-tag>
                <a-tag color="red" @click.stop="del(item.id)">删除</a-tag>
              </div>
            </div>
          </template>
        </a-tree>
      </div>
      <div class="form">
        <a-form  :form="form" style="width: 500px">
          <a-form-item label="菜单标题">
            <a-input placeholder="菜单标题"
                     v-decorator="['title', { rules: [{ required: true, message: '菜单标题' }] }]"
            />
          </a-form-item>
          <a-form-item label="权限名称">
            <a-input placeholder="权限名称"
                     v-decorator="['menu_id', { rules: [{ required: true, message: '权限名称' }] }]"
            />
          </a-form-item>
          <a-form-item v-if="parentId" label="父级菜单">
            <a-input placeholder="父级菜单" v-decorator="['parentTitle', { rules: [{ required: true, message: '菜单路径' }] }]" />
          </a-form-item>
          <a-form-item label="菜单路径">
            <a-input placeholder="菜单路径"  v-decorator="['path', { rules: [{ required: true, message: '菜单路径' }] }]" />
          </a-form-item>
          <a-form-item label="菜单名称">
            <a-input placeholder="菜单名称"   v-decorator="['name', { rules: [{ required: true, message: '菜单名称' }] }]" />
          </a-form-item>
          <a-form-item label="菜单图标">
            <IconSelector  @change="handleIconChange" v-model="currentSelectedIcon"/>
          </a-form-item>
          <a-form-item>
            <div style="text-align: center">
              <a-button @click="save">保存</a-button>
              <a-divider type="vertical" />
              <a-button @click="primaryMenu" :disabled="!parentId">新增主菜单</a-button>
            </div>
          </a-form-item>
        </a-form>
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
    }
  }
}
import IconSelector from '@/components/IconSelector'
export default {
  name: 'TableList',
  components: {
    [tree.name]: tree,
    IconSelector
  },
  data() {
    return {
      treeData: [],
      form: this.$form.createForm(this),
      parentId: '',
      parentTile: {},
      currentSelectedIcon: 'pause-circle',
    };
  },
  created() {
    this.getData()
  },
  methods: {
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

<style lang="less" scoped>
.display {
  display: flex;
  .form {
    flex: 0 0 500px;
  }
  .custom-tree {
    flex: 0 0 calc(100% - 500px);
    .title {
      align-items: center;
      display: flex;
      .label {
        margin-right: 10px;
      }
    }
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
