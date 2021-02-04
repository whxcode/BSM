<template>
  <a-modal
    :title="id !== undefined ? '修改角色': '添加角色'"
    :width="640"
    :visible="visible"
    :confirmLoading="loading"
    @ok="handleOk"
    @cancel="() => { visible = false }"
  >
    <a-spin :spinning="loading">
      <a-form :form="form">
        <a-form-item label="角色名称">
          <a-input
            v-decorator="['label', { rules: [{ required: true, message: '角色名称' }] }]"
          />
        </a-form-item>
        <a-form-item label="菜单列表">
          <a-select
            mode="multiple"
            v-decorator="[
            'actions',
            { rules: [{ required: true, message: '请选择菜单列表' }] },
          ]"
            placeholder="请选择菜单列表"
          >
            <a-select-option v-for="item in menus" :key="item.id">
                {{ item.permissionName }}
            </a-select-option>
          </a-select>
        </a-form-item>
      </a-form>
    </a-spin>
  </a-modal>
</template>

<script>
import pick from 'lodash.pick'

// 表单字段
const fields = ['description', 'id']
import { authButton } from '@/config/defaultSettings'
import { queryMenusList,roleGetQueryMenus,roleAdd  } from '@/api/manage'

export default {
  props: {
    model: {
      type: Object,
      default: () => null
    }
  },
  data() {
    return {
      loading: false,
      id: null,
      visible: false,
      form: this.$form.createForm(this),
      operator: authButton,
      menus: []
    }
  },
  created() {
    // 防止表单未注册
    fields.forEach(v => this.form.getFieldDecorator(v))
    queryMenusList({ pageSize: 1000,pageNo: 1 }).then(({ result: { data } }) => {
      this.menus = data
      console.log(this.menus)
    })
    // 当 model 发生改变时，为表单设置值
    this.$watch('model', () => {
      this.model && this.form.setFieldsValue(pick(this.model, fields))
    })
  },
  methods: {
    handleOk () {
      const form = this.form
      form.validateFields(async (errors, values) => {
        if(!errors) {
          if(this.id) {
            values.id = this.id
          }
          this.loading = true

          await roleAdd(values)
          this.loading = false
          this.visible = false
          this.form.resetFields()
          this.id = ''

         this.$emit('ok')
          this
        } else {

        }
      })
    },
    async init(roleId) {
      this.visible = true
      this.loading = true
      const data = await roleGetQueryMenus({ roleId })
      console.warn(data)
      const actions = data.menus.map(item => item.id)
      this.form.setFieldsInitialValue({ label: data.label,actions })
      this.id = roleId
      this.loading = false
    }
  }
}
</script>
