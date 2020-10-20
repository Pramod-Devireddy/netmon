/*
 * File         : LogDataPage.vue
 * Project      : NetMon
 * Created Date : Wednesday, Oct 7th 2020, 2:37:13 AM
 * Author       : Pramod Devireddy
 * 
 * Last Modified: Wednesday, 7th October 2020 9:09:29 am
 * Modified By  : Pramod Devireddy
 * 
 * Copyright (c)2020 - Pramod Devireddy
 * ************************* Description *************************
 *  
 * ***************************************************************
 */

<template>
  <div>
    <div class="d-flex justify-center" style="border-radius: 5px; background-color: #333; padding: 10px">
      <div id="myDiv" height="100%"></div>
    </div>
  <div>{{getRxMax}} kBits/s</div>
  <div>{{getTxMax}} kBits/s</div>
  </div>
</template>


<script>
module.exports = {
  name: "Plotly",
  data() {
    return {
      logData: {},
      trace1: {
        x: [],
        y: [],
        name: "Rx Traffic",
        fill: "tozeroy",
        type: "scattergl",
      },
      trace2: {
        x: [],
        y: [],
        name: "Tx Traffic",
        fill: "tozeroy",
        type: "scattergl",
      },
    };
  },
  methods: {},
  mounted: function () {
    var vm = this;

    axios
      .get("http://192.168.1.104:9989/rx_log_data")
      .then(function (response) {
        vm.trace1.x = response.data.time;
        vm.trace1.y = response.data.bytes;
      })
      .catch(function (error) {
        console.log(error);
      });

    axios
      .get("http://192.168.1.104:9989/tx_log_data")
      .then(function (response) {
        vm.trace2.x = response.data.time;
        vm.trace2.y = response.data.bytes;
      })
      .catch(function (error) {
        console.log(error);
      });

    var layout = {
      uirevision: "true",
      title: "Network Traffic",
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

    Plotly.newPlot("myDiv", [this.trace1, this.trace2], layout, {
      displaylogo: false,
      responsive: true,
      scrollZoom: true,
    });
  },
  watch: {
    trace1: {
      handler: function () {
        var layout = {
          legend: {
            font: {
              size: 12,
              color: "#FFF",
            },
          },

          margin: {
            l: 120,
            r: 20,
            b: 75,
            t: 50,
            pad: 5,
          },
          width: "1200",
          height: "800",
          uirevision: "true",
          title: "Network Traffic",
          titlefont: {
            size: 18,
            color: "#FFF",
          },
          paper_bgcolor: "#333",
          plot_bgcolor: "#333",
          xaxis: {
            color: "#FFF",
            showline: true,
            linewidth: 2,
            tickfont: {
              size: 16,
            },
          },
          yaxis: {
            color: "#FFF",
            nticks: 10,
            showline: true,
            zeroline: false,
            linecolor: "#fff",
            linewidth: 2,
            tickfont: {
              size: 16,
            },

            title: "Traffic",
            titlefont: {
              size: 18,
              color: "#FFF",
            },
            tickformat: " ",
            ticksuffix: " kb/s",
          },
        };
        layout.datarevision = new Date().getTime();
        Plotly.react("myDiv", [this.trace1, this.trace2], layout);
      },
      deep: true,
    },
    trace2: {
      handler: function () {
        var layout = {
          legend: {
            font: {
              size: 12,
              color: "#FFF",
            },
          },
          margin: {
            l: 120,
            r: 20,
            b: 75,
            t: 50,
            pad: 5,
          },
          width: "1200",
          height: "800",
          uirevision: "true",
          title: "Network Traffic",
          titlefont: {
            size: 18,
            color: "#FFF",
          },
          paper_bgcolor: "#333",
          plot_bgcolor: "#333",
          xaxis: {
            color: "#FFF",
            showline: true,
            linewidth: 2,
            tickfont: {
              size: 16,
            },
          },
          yaxis: {
            color: "#FFF",
            nticks: 10,
            showline: true,
            zeroline: false,
            linecolor: "#fff",
            linewidth: 2,
            tickfont: {
              size: 16,
            },

            title: "Traffic",
            titlefont: {
              size: 18,
              color: "#FFF",
            },
            tickformat: " ",
            ticksuffix: " kb/s",
          },
        };
        layout.datarevision = new Date().getTime();
        Plotly.react("myDiv", [this.trace1, this.trace2], layout);
      },
      deep: true,
    },
  },
  computed: {
    getTxMax() {
      return Math.max.apply(null, this.trace2.y)
    },
    getRxMax() {
      return Math.max.apply(null, this.trace1.y)
    },
  }
};
</script>
