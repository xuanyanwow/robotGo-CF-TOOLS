# RobotGo-FPS-TOOLS
辅助脚本

> 仅作为自动化测试的学习项目，不得用于游戏作弊，破坏游戏平衡；精力有限将加入数值的阈值设置 如禁止低延迟连续操作

- [x] F1-USP连点器
- [ ] F2-M4远距连点
- [ ] F3-加特林连点
- [ ] F4-自动瞄准(robotgo支持图像搜索后)
- [x] F5-视频笔记切换


## USP

```shell
# 70~120延迟 连发3~4发
main.exe -minSleep=70 -maxSleep=120 -minShut=3 -maxShut=4
```