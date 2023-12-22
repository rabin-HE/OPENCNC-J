![](/Picture/machine.jpg)

## **OPENCNC-J**  [中文README](/README.MD)

* Introduction
  * This is an open-source integration project based on Linuxcnc. Its aim is to facilitate electrical engineers in using a combination of Linuxcnc + EtherCAT with CIA402 protocol + Raspberry Pi (or X86 - Debian).
* Current Version - Looking forward to more developers joining
  * 0.1.0
* Personal Video Release Page - Detailed Project Content
  * [Bilibili - OPENCNC-J Development Process Collection](https://space.bilibili.com/341589947/channel/collectiondetail?sid=1918828)
* Notice & Warning!!!
  * This is an open-source project. It does not provide any warranty or compensation, and involves electrical safety warnings. Please be responsible for it yourself.
  * This project follows the GPL-3.0 open-source license.

## **Installation Instructions - Taking Raspberry Pi Installation as an Example**

Install [Linuxcnc](https://linuxcnc.org/)

1. Download the official installation package - LinuxCNC 2.9.1[Raspberry Pi OS based on Debian 12 Bookworm](https://www.linuxcnc.org/iso/linuxcnc-2.9.1-bookworm-rpi4.img.xz)
2. Create an SD card system - It is recommended to use the[official Raspberry Pi](https://www.raspberrypi.com/software/https:/) SD card system creation tools for[Windows](https://downloads.raspberrypi.org/imager/imager_latest.exehttps:/) &[Ubuntu](https://downloads.raspberrypi.org/imager/imager_latest_amd64.debhttps:/)
3. Insert the card to boot the Raspberry Pi
4. Configure Linuxcnc - [To be filled in]

Install dependencies - [Refer to the EtherCAT installation example](https://forum.linuxcnc.org/ethercat/45336-ethercat-installation-from-repositories-how-to-step-by-stephttps:/), CIA402

1. EtherCAT installation
   1. Open terminal:`sudo apt install ethercat-master libethercat-dev linuxcnc-ethercat`
   2.
   3. Check the MAC address of the network port:`ip a`
      1. Record the MAC address, e.g.,`xx:aa:yy:zz:bb:cc`
   4. Enter (configure) into the ethercat.conf file
      1. `sudo geany /etc/ethercat.conf`
      2. Modify the file with geany
         1. `MASTER0_DEVICE="xx:aa:yy:zz:bb:cc"` - Replace with the MAC address you just obtained`DEVICE_MODULES="generic"`
   5. Set the EtherCAT service to start on boot
      1. `sudo systemctl enable ethercat.service`
      2. `sudo systemctl start ethercat.service`
      3. `sudo systemctl status ethercat.service`
      4. `sudo chmod 666 /dev/EtherCAT0`
      5. EtherCAT test command
         1. `ethercat slaves` - Check for errors to confirm successful installation
         2. Restart
2. Install CIA402
   1. Return to the user directory:`cd`
   2. `git clone https://github.com/dbraun1981/hal-cia402` - If there's an error, please confirm if git tool is installed
      1. `sudo apt install git` - Git installation command
   3. Enter the hal-cia402 folder:`cd hal-cia402`
   4. Compile and install the 402 tool package:`sudo halcompile --install cia402.comp`

That completes the most challenging part 👍

## **Driver Parameter Settings**

* When starting, shows AL221 alarm [Encoder not zeroed]
  * Set zero point: Click on driver panel operation → Switch to [AF.CEN] → Press S → Press S again → Set successfully #Refer to the driver manual P39 - 4.3.3 Absolute operation
* Motor Reversing
  * If the motor direction is incorrect, the motor direction can be switched
    * Setting method: Modify parameter P1.01 = 1 or 0
* Electronic Gear Ratio
  * Electronic gear ratio numerator P3.08 Electronic gear ratio denominator P3.10

## **Directory Structure**

* CAD - Electrical schematic diagram in PDF format
* linuxcnc - Contains all the required configuration files

## **Project Procurement List**

Disclaimer: There is no commercial relationship with the procurement links, please inquire and compare prices yourself.


| Item                        | URL                                                                                                                                                                             | Model                                                                           | Quantity            | Notes                                                                                                |
| ----------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------- | ------------------------------------------------------------------------------------------------------ |
| XY Axis Motor+Drive         | [URL](https://detail.tmall.com/item.htm?id=651348933042&spm=2015.23436601.0.0)                                                                                                  | (New) 400W Servo+Drive+5m Cable<br/>Motor Model: RSM-M06L1330A (No Brake)       | 2                   | Multi-turn Encoder<br/>Absolute Origin!!! This costs extra<br/>5m Cable<br/>Supports EtherCAT+CIA402 |
| Raspberry Pi_4 Dev Board    | [URL](https://detail.tmall.com/item.htm?abbucket=2&id=608798378397&ns=1&skuId=4551930747475&spm=a230r.1.14.43.2fad28eea6m3ZB)                                                   | Pi 4B/8G[Essential Package]                                                     | 1                   |                                                                                                      |
| Micro HDMI to HDMI Cable    | [URL](https://detail.tmall.com/item.htm?abbucket=2&id=658377041318&ns=1&skuId=4747932408936&spm=a230r.1.14.7.399d1c02A92EDD)                                                    | 3m                                                                              | 1                   |                                                                                                      |
| EtherCAT Servo Bus          | [URL](https://detail.tmall.com/item.htm?abbucket=2&id=674202033100&ns=1&sku_properties=1627207:21423176984&spm=a230r.1.14.1.4e83345adWD5Zj)                                     | High Flexibility Gigabit Shielded Ethernet Cable - Green                        | 3(0.5m)<br/>1(1.5m) | -                                                                                                    |
| Z Axis Motor+Drive          | [URL](https://detail.tmall.com/item.htm?id=651348933042&spm=2015.23436601.0.0&skuId=4870002170807)                                                                              | (New) RS400E Servo+Drive+5m Cable<br/>Motor Model: RSM-M06L1330A-Z (With Brake) | 1                   | Multi-turn Encoder<br/>Absolute Origin!!! This costs extra<br/>5m Cable<br/>Supports EtherCAT+CIA402 |
| Spindle Motor+Drive         | [URL](https://detail.tmall.com/item.htm?id=651348933042&spm=2015.23436601.0.0&skuId=4870002170807)                                                                              | (New) RS400E Servo+Drive+5m Cable<br/>Motor Model: RSM-M06J1330A (No Brake)     | 1                   | Standard Encoder<br/>5m Cable<br/>Supports EtherCAT+CIA402                                           |
| Raspberry Pi_Rail Box       | [URL](https://item.taobao.com/item.htm?spm=a230r.1.14.7.ae487264WDE38q&id=663400858369&ns=1&abbucket=2#detail)                                                                  |                                                                                 | 1                   |                                                                                                      |
| Monitor with Case           | [URL](https://detail.tmall.com/item.htm?id=666450227637&skuId=5093251227751&spm=pc_detail.27183998.202208.1.705a7dd62FFPiD)                                                     | 7 inch 1024*600 IPS+Touch+Case                                                  | 1                   | Not recommended,<br/>Broke in less than a month                                                      |
| Monitor Stand               | [URL](https://detail.tmall.com/item.htm?abbucket=17&id=687974425740&rn=c8d5c0a97e71825d8e11d5a7d7fd0f17&spm=a1z10.5-b.w4011-23875343074.107.14993b2bXRLXEy&skuId=4901270192689) |                                                                                 | 1                   |                                                                                                      |
| Handwheel                   | [URL](https://item.taobao.com/item.htm?spm=a21n57.1.0.0.10a6523cQWLmki&id=45130365561&ns=1&abbucket=0#detail)                                                                   | MACH3 Wired Handwheel 4 Axis                                                    | 1                   |                                                                                                      |
| Rail Socket                 | [URL](https://detail.tmall.com/item.htm?abbucket=2&id=677701080516&ns=1&spm=a21n57.1.0.0.162c523chCpKY3&skuId=5038263790059)                                                    | 5 Slots                                                                         | 1                   |                                                                                                      |
| USB Cable                   | [URL](https://item.taobao.com/item.htm?spm=a21n57.1.0.0.19d6523c0w91YJ&id=542169862221&ns=1&abbucket=2#detail)                                                                  | Micro USB Data Cable Black 5m                                                   | 1                   |                                                                                                      |
| Axial Fan                   | [URL](https://detail.tmall.com/item.htm?abbucket=2&id=658927980547&ns=1&skuId=4759805705081&spm=a21n57.1.0.0.6a79523cvWfTb0)                                                    | Flame Retardant ZL-803+12038 Fan AC220V+Protective Net                          | 2                   |                                                                                                      |
| Emergency Stop Button       | [URL](https://detail.tmall.com/item.htm?abbucket=2&id=552913814211&ns=1&spm=a21n57.1.0.0.347b523ctQ5vgw)                                                                        |                                                                                 |                     |                                                                                                      |
| Circuit Breaker (Schneider) |                                                                                                                                                                                 | IC65N-C10A/2P 10A<br/>IC65N-C6A/2P 6A                                           | 1<br/>2             | Please purchase on your own                                                                          |
| 24V DC Power Supply         |                                                                                                                                                                                 | MW Power Supply, 120W                                                           | 1                   | Please purchase on your own                                                                          |

## Special Thanks

* Thanks to Guilin Guanglu Digital Measurement and Control Co., Ltd. for their support of this open-source project
  * Tmall Flagship Store -[Product Purchase Address](https://guanglujj.tmall.com/shop/view_shop.htm?spm=a21n57.1.0.0.396d523cnfFegq&appUid=RAzN8HWJMBXJctTqz11nosKY98Wm2AwEqKJEpqoPs9YQHKEVNDc)![Guilin Guanglu](https://chat.openai.com/Picture/guilinguanglu.png)
* Thanks to the developers who helped me during the development process[Hakan](https://forum.linuxcnc.org/cb-profile/22448-hakan),[rodw](https://forum.linuxcnc.org/cb-profile/rodw), and[garlicbread](https://forum.linuxcnc.org/cb-profile/garlicbread)
* Project References
  * [YouTube - Ethercat + linuxcnc + Raspberry pi?](https://www.youtube.com/watch?v=NQ-HnrusGJo&t=16s)
