<template>
  <div class="account-settings-info-view">
    <a-row :gutter="16">
      <a-col :md="24" :lg="16">
        <a-form layout="vertical" :form="form">
          <a-form-item
            :label="$t('account.settings.basic.nickname')"
          >
            <a-input :placeholder="$t('account.settings.basic.nickname-message')"
                     v-decorator="['surname', { rules: [{ required: true, message: '用户名称' }] }]"
            />
          </a-form-item>
          <a-form-item
            :label="$t('account.settings.basic.profile')"
          >
            <a-textarea rows="4" :placeholder="$t('account.settings.basic.profile-message')"
                        v-decorator="['brief', { rules: [{ required: true, message: '个人简介' }] }]"
            />
          </a-form-item>

          <a-form-item
            :label="$t('account.settings.basic.email')"
            :required="false"
          >
            <a-input placeholder="example@ant.design"   v-decorator="['email', { rules: [{ required: true, message: '个人邮箱' }] }]"/>
          </a-form-item>

          <a-form-item>
            <a-button type="primary" @click="submit">{{ $t('account.settings.basic.update') }}</a-button>
          </a-form-item>
        </a-form>
      </a-col>
      <a-col :md="24" :lg="8" :style="{ minHeight: '180px' }">
        <div class="ant-upload-preview" @click="$refs.modal.edit(1)" >
          <a-icon type="cloud-upload-o" class="upload-icon"/>
          <div class="mask">
            <a-icon type="plus" />
          </div>
          <img :src="option.img"/>
        </div>
      </a-col>
    </a-row>
    <avatar-modal ref="modal" @ok="setavatar"/>
  </div>
</template>

<script>
import AvatarModal from './AvatarModal'
import { modifyBaseUser } from '@/api/manage'

export default {
  components: {
    AvatarModal
  },
  data () {
    return {
      form: this.$form.createForm(this),
      preview: {},
      user: {},
      option: {
        img: '/avatar2.jpg',
        info: true,
        size: 1,
        outputType: 'jpeg',
        canScale: false,
        autoCrop: true,
        // 只有自动截图开启 宽度高度才生效
        autoCropWidth: 180,
        autoCropHeight: 180,
        fixedBox: true,
        // 开启宽度和高度比例
        fixed: true,
        fixedNumber: [1, 1]
      }
    }
  },
  methods: {
    setavatar (url) {
      this.option.img = url
    },
    submit() {
      this.form.validateFields(async (err,values) => {
        if(err) {
          return
        }
        values.avatar = this.option.img
        values.id = this.user.userInfo.id
        await modifyBaseUser(values)

        const data = {
          ...this.user,
          userInfo: {
            ...values,
          }
        }
        this.$store.commit('SET_REAL_USER',data)
        this.$ls.set("userInfo",data)
      })
    }
  },
  mounted() {
    this.$nextTick(() => {
      const user = this.$ls.get('userInfo')
      this.user = user
      this.form.setFieldsValue(user.userInfo)
      this.option.img = user.userInfo.avatar
    })

  }
}
</script>

<style lang="less" scoped>

  .avatar-upload-wrapper {
    height: 200px;
    width: 100%;
  }

  .ant-upload-preview {
    position: relative;
    margin: 0 auto;
    width: 100%;
    max-width: 180px;
    border-radius: 50%;
    box-shadow: 0 0 4px #ccc;

    .upload-icon {
      position: absolute;
      top: 0;
      right: 10px;
      font-size: 1.4rem;
      padding: 0.5rem;
      background: rgba(222, 221, 221, 0.7);
      border-radius: 50%;
      border: 1px solid rgba(0, 0, 0, 0.2);
    }
    .mask {
      opacity: 0;
      position: absolute;
      background: rgba(0,0,0,0.4);
      cursor: pointer;
      transition: opacity 0.4s;

      &:hover {
        opacity: 1;
      }

      i {
        font-size: 2rem;
        position: absolute;
        top: 50%;
        left: 50%;
        margin-left: -1rem;
        margin-top: -1rem;
        color: #d6d6d6;
      }
    }

    img, .mask {
      width: 100%;
      max-width: 180px;
      height: 100%;
      border-radius: 50%;
      overflow: hidden;
    }
  }
</style>
