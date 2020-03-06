<template lang="pug">
  div.app
    zi-image(
      v-if="!success"
      width="40%"
      v-bind:src="imagepng"
    )
    zi-row(justify="center")
      h1 {{ text }}
    zi-row(justify="center" style="margin-top: 10px;")
      zi-col(span="8")
        zi-input(placeholder="Input your search key" v-model="searchKey" style="width:100%;" size="big")
    zi-row(justify="center" style="margin-top: 10px;")
      zi-checkbox(v-model="confirm") 请在点击搜索前确保已经阅读并知晓相关法律，并为自己的行为负责
    zi-row(justify="center" style="margin-top: 20px;")
      zi-button(type="primary" @click="onSearchButtonClick") search

    zi-row(v-if="success" justify="center" style="margin-top: 30px;")
      zi-col(span="14")
        zi-card()
          zi-row(justify="center")
            h3 Search Result
          zi-row
            zi-col(span="20")
              zi-table(:data="result" hover width="1000" style="margin-left: 30px;")
                zi-table-column(prop="hit" label="Hit" width="150")
                zi-table-column(prop="text" label="Content")
          zi-more(expand text="付费查看更多")
</template>


<script>

  export default {
  name: 'App',
  components: {},
  data: () => ({
    imagepng: require('./assets/logo.jpg'),
    text: 'The most powerful social work tool',
    searchKey: '',
    confirm: true,
    result: [{"score":1,"hit":"山东大学","text":"15391,18253160182,di720805,hui720805,狄辉辉,372924199108051819,1075173797@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"19229,18769785165,gua0305,20070305,张晨,460102199203031224,1187452092@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"35411,18363127029,zhuyuyouka,zhuyu931010,朱煜,372923199412040069,752320082@qq.com,您的大学校名是？,山东大学威海分校"},{"score":1,"hit":"山东大学","text":"35881,18253162967,k991082295,shan160dong,洪克勤,350583199303231017,991082295@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"51887,18363125356,ghnghn001,ghnh2co3,龚浩宁,320602199206271510,120727231@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"94749,18311091836,klnhbin,zhoubin123,周彬,412826199304203533,614549942@qq.com,您的大学校名是？,山东大学威海分校"},{"score":1,"hit":"山东大学","text":"98009,18769787739,blx19911128,19911128blx,刘栋,370781199111280270,806013716@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"102890,15901856956,du9fu9,87208728dufu,杜福,370281198809270515,du9fu9@yahoo.com.cn,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"121818,18766317478,n1256253432,n249705277,牛宝东,372324199110130011,1256253432@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"123600,18253163821,lxb1539376074,lxb19920115,李茜博,412321199201150011,1539376074@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"126724,18253166137,wanglinqi0216,qi19940216,王琳琪,370724199402167289,2435050921@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"147191,13573748724,gmy20002007,gmy20002007,龚梦尧,310115199208142715,624804366@qq.com,您的大学校名是？,山东大学威海分校"},{"score":1,"hit":"山东大学","text":"153647,13918377361,oldwanglaowu,aaaaaaa8,郑立,310115199012307813,844154950@qq.com,您的出生地是？,山东大学"},{"score":1,"hit":"山东大学","text":"159565,18661461581,lxcctv4310,lx824310,李鑫,370203198212097618,lxcctv4310@163.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"175977,15866721630,delinkeye,950523lzy,李宗尧,340204199505232014,delink@163.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"194534,18910918815,livelyxin,gengjia123,齐文光,372330198310123759,4332169@163.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"197131,15953909698,diamondsuper,831001,马俊月,372526198310010726,diamondsuper@163.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"222791,18231655035,heruicome,heruicome0622,贺蕊,130632198906074428,hehau@163.com,您的大学校名是？,山东大学威海分校"},{"score":1,"hit":"山东大学","text":"239157,13561125632,rockdj,fuq_314159,傅强,371428198303191035,yaogunheike@163.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"249840,15265356927,xixishou,wang11,王茜,371081198906192424,578788229@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"257789,18716437227,dazhi5190,370784,刘治涛,370784198403075011,285879075@qq.com,您的大学校名是？,山东大学威海分校"},{"score":1,"hit":"山东大学","text":"258986,13912604614,yishengyouni922,cheng870922,程玉洁,622301198611023321,296551375@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"310790,18653475552,hon_wu,2673hw,洪德刚,37240119790302371X,254141241@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"326678,15005314287,kira2012,llc1234,刘鲁超,372922199101243073,lucharo@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"328146,18463105978,qwer660630,zxp13679,李响,370883199601285515,542003683@qq.com,您的大学校名是？,山东大学威海分校"},{"score":1,"hit":"山东大学","text":"328376,15829085962,surikun,sk20091016,董素珍,411081195404019225,276906571@qq.com,您的大学校名是？,山东大学额"},{"score":1,"hit":"山东大学","text":"331020,13465237661,thankgod,mm.777,娄娟,370125199007147468,609039891@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"337790,13553197896,renpingli,rpl680504,任平礼,370687198510020693,393203227@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"345982,15562107257,limenggege,xiao88xiao88,李梦歌,410527199207072023,1154136151@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"349654,15069117392,lgq123l,lgq123l,李国强,371312198908105551,994394097@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"368213,15614667720,wgy20980,yrai1314,魏光煜,341182199310201634,404120980@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"371687,13965770460,xiaobin121,xia8812784,肖彬,370921198710100243,860626131@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"377611,18769781611,qxliuxiaonan,lnx181721,刘晓楠,370686199008245510,852947641@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"379560,15098853630,woshirzf,q1w2e3r4,任中飞,370102199212264511,1015464289@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"392987,18297001003,QQ903351084,19911118mst,马舒婷,632124199111180021,903351084@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"399996,15910093956,quj0711,qj19910526,曲杰,370612199107073510,351593512@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"401394,13585896526,siyecaoyi,jimusaer,孙怡,320483198809210920,742972558@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"406760,18660816891,elma201208,elma273,李侠,370402198305270023,284984673@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"410566,18364997216,tiefuqiang,123nie,铁福强,632122199010085014,feigerui@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"420383,18055275570,ahwangzd,wzd880522,王宗鼎,340321198805220178,253630867@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"435103,15865350561,laixilihao,woshilihao,李浩,370285198801012912,404904188@qq.com,您的大学校名是？,山东大学"},{"score":1,"hit":"山东大学","text":"446596,15265623789,jkmoka,jzm521,鞠钊民,370782198801302416,273478115@qq.com,您的大学校名是？,山东大学"}],
    success: false,
    host: "127.0.0.1:19090",
    agreement: "http:",
  }),
  created() {
    this.host = location.host
    this.agreement = location.protocol
  },
  methods: {
    onSearchButtonClick() {
      var that = this
      if (this.searchKey==="") {
        const tt = that.$Toast.show({
          action: 'close',
          text: '搜索关键词请不要留空',
          handler: () => tt.close(),
        })
      } else  {
        that.axios.get(that.agreement+"//"+that.host+"/search?key="+that.searchKey).then((response) => {
          if (response.status===200) {
            console.log(response)
            if (response.data==null) {
              that.selectError()
            } else {
              that.result = response.data
              that.success = true
            }
          } else {
            that.selectError()
          }
        }).catch(() => {
          that.selectError()
        })
      }
    },
    selectError() {
        const instance = this.$Toast.show({
          type: 'danger',
          action: 'close',
          text: '查询失败',
          handler: () => instance.close(),
        })
      this.success = false
    }
  }
}
</script>

<style>

</style>
