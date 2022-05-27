# dorm2free

我在宿舍的远程办公+免流量计费方案的配套工具 for linux dekstop

相关配置文件 + 一个系统托盘 GUI

# 方案

见[这里](https://blog.zilch40.wang/post/my-experience-on-ipv6-mianliu/)

# 安装依赖

Archlinux with archlinux-cn 源

```
sudo pacman -S libayatana-appindicator clash-premium-bin
```

# files

- 将 `files` 下的 clash 配置文件和 systemd service 放在 `~/.config`
- 用 `scripts` 下的脚本切换访问模式：
    - clash
        - `rule` 全部 TCP 流量经过 pg 出去
        - `direct` 全部 TCP 流量直连，不经过 pg 
    - pg
        - `v2ray` 全部 TCP 流量经 pg 从远端的 v2ray 按 IP 分流出去
        - `direct` 全部 TCP 流量经 pg 从远端直连出去 (访问 IEEE Xplore 等需要 懒得写规则)

# GUI \[optional\]

- `build.sh` 编译
- 常驻系统托盘, 鼠标切换模式

# 问题

1. 使 ping 结果不准确
2. 没考虑 UDP
3. 考虑到安全隐患，不推荐在单位自建服务
