<template>
  <a-modal
    title="新建菜单"
    :width="640"
    :visible="visible"
    :confirmLoading="loading"
    @ok="handleOk"
    @cancel="() => { visible = false }"
  >
    <a-spin :spinning="loading">
      <a-form :form="form">
        <a-form-item label="菜单名称">
          <a-input
            v-decorator="['permissionName', { rules: [{ required: true, message: '菜单名称' }] }]"
          />
        </a-form-item>
        <a-form-item label="菜单权限Id">
          <a-input
            v-decorator="['permissionId', { rules: [{ required: true, message: '菜单权限名称' }] }]"
          />
        </a-form-item>
        <a-form-item label="功能权限按钮">
          <a-select
            mode="multiple"
            v-decorator="[
            'actions',
            { rules: [{ required: true, message: '功能权限' }] },
          ]"
            placeholder="菜单权限Id"
          >
            <a-select-option v-for="item in operator" :key="item.action">
                {{ item.describe }}
            </a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="菜单图标">
          <a-input
            v-decorator="['icon', { rules: [{ required: true, message: '菜单图标' }] }]"
          />
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
import { menuAdd, menuGet } from '@/api/manage'

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
      operator: authButton
    }
  },
  created() {
    // 防止表单未注册
    fields.forEach(v => this.form.getFieldDecorator(v))

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
          await menuAdd(values)
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
    async init(menuId) {
      this.visible = true
      this.loading = true
      const data = await menuGet({ menuId })
      data.actions = JSON.parse(data.actions)
      this.form.setFieldsValue({ ...data })
      this.id = menuId
      this.loading = false
    }
  }
}
</script>
