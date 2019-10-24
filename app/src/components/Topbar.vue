<template>
  <el-header>
    <el-row>
      <el-col :span="12">
        <el-page-header @back="goBack"
                        :content="title"></el-page-header>
      </el-col>
      <el-col :span="12" class="align-right">
        <div class="vertical-center">
          <i class="el-icon-bell"></i>
          <el-divider direction="vertical"></el-divider>
          <el-avatar shape="circle"
                     :size="32"
                     src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"></el-avatar>
        </div>
      </el-col>
    </el-row>
  </el-header>
</template>
<script>
  export default {
    data() {
      return {}
    },
    props: ["title"],
    created() {
      document.title = this.title;
    },
    methods: {
      goBack() {
        history.back();
      }
    }
  };
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

  .el-header {
    background-color: #fff;
  }
</style>
