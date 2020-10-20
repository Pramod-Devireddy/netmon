/*
 * File         : plotly.vue
 * Project      : NetMon
 * Created Date : Wednesday, Oct 7th 2020, 1:20:14 AM
 * Author       : Pramod Devireddy
 * 
 * Last Modified: Wednesday, 7th October 2020 3:04:14 am
 * Modified By  : Pramod Devireddy
 * 
 * Copyright (c)2020 - Pramod Devireddy
 * ************************* Description *************************
 *  
 * ***************************************************************
 */


<template>
  <div>
    <div id="myDiv" height="100%"></div>
  </div>
</template>

<script>
module.exports = {
  name: "Plotly",
  data() {
    return {
      trace1: {
        x: [],
        y: [],
        type: "scattergl",
      },
    };
  },
  methods: {},
  mounted: function () {
    var vm = this;
    var tm1WS = new WebSocket("ws://127.0.0.1:9050/ws");
    tm1WS.onopen = function () {
      tm1WS.send('{"action":"subscribe","params":["Frame-Id"]}');
    };
    tm1WS.onmessage = (msg) => {
      resp = JSON.parse(msg.data);
      this.trace1.x.push(
        moment(resp.time_stamp, "DD-MMM-YYYY HH:mm:ss:SSS").format(
          "YYYY-MM-DD HH:mm:ss.SSS"
        )
      );
      this.trace1.y.push(resp.proc_value);
      // Plotly.extendTraces("myDiv", { y: [[resp.proc_value]] }, [0]);
      console.log(resp.proc_value);
    };

    var layout = {
      width: 1000,
      height: 800,
      uirevision: "true",
      title: "Line and Scatter Plot",
      paper_bgcolor: "#333",
      plot_bgcolor: "#333",
      xaxis: {
        color: "#FFF",
        tickfont: {
          size: 14,
        },
      },
      yaxis: {
        color: "#FFF",
        showline: true,
        linecolor: "#fff",
        zeroline: true,
        linewidth: 2,
        tickfont: {
          size: 14,
        },
      },
    };

    Plotly.newPlot("myDiv", [this.trace1], layout, {
      displaylogo: false,
      responsive: true,
      scrollZoom: true,
    });
  },
  watch: {
    trace1: {
      handler: function () {
        var layout = {
          width: 1000,
          height: 800,
          uirevision: "true",
          title: "Line and Scatter Plot",
          paper_bgcolor: "#333",
          plot_bgcolor: "#333",
          xaxis: {
            color: "#FFF",
            tickfont: {
              size: 14,
            },
          },
          yaxis: {
            color: "#FFF",
            nticks: 10,
            showline: true,
            zeroline: true,
            linecolor: "#fff",
            linewidth: 2,
            tickfont: {
              size: 14,
            },
          },
        };
        layout.datarevision = new Date().getTime();
        Plotly.react("myDiv", [this.trace1], layout);
      },
      deep: true,
    },
  },
};
</script>

<style scoped>
</style>