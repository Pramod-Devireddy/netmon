/*
 * File         : rx_netmon.sh
 * Project      : NetMon
 * Created Date : Tuesday, Oct 20th 2020, 6:48:02 PM
 * Author       : Pramod Devireddy
 * 
 * Last Modified: Tuesday, 20th October 2020 6:48:12 pm
 * Modified By  : Pramod Devireddy
 * 
 * Copyright (c)2020 - Pramod Devireddy
 * ************************* Description *************************
 *  
 * ***************************************************************
 */


cd /sys/class/net/enp0s3/statistics

echo `date`

old="$(<rx_bytes)"; while $(sleep 1); do
now=$(<rx_bytes); echo $((($now-$old)*8/1000)) kBit/s; old=$now; done
