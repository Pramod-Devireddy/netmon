new Vue({
  el: "#app",
  vuetify: new Vuetify(),
  data: {
    mw1speed: "0",
    mw2speed: "0",
    tmwspeed: "0",
    mw1cur: "0",
    mw2cur: "0",
    tmwcur: "0",
    mw1sts: "",
    mw2sts: "",
    tmwsts: "",
    wde1sts: "",
    wde2sts: "",
    wde3sts: "",
    mw1bearingtmp: "0",
    mw2bearingtmp: "0",
    rwbearingtmp: "0",
    mw1basetmp: "0",
    mw2basetmp: "0",
    rwbasetmp: "0",
    tmwrotdir: "",
  },
  components: {},
  created: function () {
    var vm = this;
    var tm1WS = new WebSocket("ws://100.200.30.5:9050/ws");
    tm1WS.onopen = function () {
      tm1WS.send('{"action":"subscribe","params":[' +
        '"mw1speed", "mw2speed", "tmwspeed",' +
        '"mw1cmdspeed", "mw2cmdspeed", "tmwcmdspeed",' +
        '"mw1cur", "mw2cur", "tmwcur",' +
        ' "mw1sts", "mw2sts", "tmwsts",' +
        '"wde1sts", "wde2sts", "wde3sts",' +
        '"mw1bearingtmp", "mw2bearingtmp", "rwbearingtmp",' +
        '"mw1basetmp", "mw2basetmp", "rwbasetmp",' +
        ' "wheelmode", "aoceselected", "aocsmode", "TMW-RotDir-AOC"]}')
    }
    tm1WS.onmessage = msg => {
      resp = JSON.parse(msg.data);

      switch (resp.param) {
        case "MW1-Sts":
          vm.mw1sts = resp.proc_value;
          break;
        case "MW2-Sts":
          vm.mw2sts = resp.proc_value;
          break;
        case "TMW-Sts":
          vm.tmwsts = resp.proc_value;
          break;
        case "WDE1-Sts":
          vm.wde1sts = resp.proc_value;
          break;
        case "WDE2-Sts":
          vm.wde2sts = resp.proc_value;
          break;
        case "WDE3-Sts":
          vm.wde3sts = resp.proc_value;
          break;
        case "MW1-Speed":
          vm.mw1speed = resp.proc_value;
          break;
        case "MW2-Speed":
          vm.mw2speed = resp.proc_value;
          break;
        case "TMW-Speed":
          vm.tmwspeed = resp.proc_value;
          break;
        case "MW1-Cur":
          vm.mw1cur = resp.proc_value;
          break;
        case "MW2-Cur":
          vm.mw2cur = resp.proc_value;
          break;
        case "TMW-Cur":
          vm.tmwcur = resp.proc_value;
          break;
        case "MW1-Bearing-Tmp":
          vm.mw1bearingtmp = resp.proc_value;
          break;
        case "MW2-Bearing-Tmp":
          vm.mw2bearingtmp = resp.proc_value;
          break;
        case "RW-Bearing-Tmp":
          vm.rwbearingtmp = resp.proc_value;
          break;
        case "MW1-Base-Tmp":
          vm.mw1basetmp = resp.proc_value;
          break;
        case "MW2-Base-Tmp":
          vm.mw2basetmp = resp.proc_value;
          break;
        case "RW-Base-Tmp":
          vm.rwbasetmp = resp.proc_value;
          break;
        case "TMW-RotDir-AOC":
          vm.tmwrotdir = resp.proc_value.trim();
          break;
      }
    }
  },
  methods: {
    getValidSpeed: function (speed) {
      if (isNaN(speed))
        return "0";
      else return speed;
    },

    getAnimationSpeed: function (speed) {
      console.log(speed)
      if (isNaN(speed))
        return 0;
      else {
        newspeed = 600 / Number(speed - 648);
        if (newspeed < 0)
          return 0;
        else
          return newspeed;
      }
    },
  },


});