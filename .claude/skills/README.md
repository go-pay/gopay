# GoPay 项目 Skills

本目录包含 GoPay 项目专用的 Claude Code skills。

## 可用 Skills

### add-payment-interface

为 GoPay 项目添加新支付接口的标准化工作流程。

**功能：**
- 分析官方接口文档
- 实现接口代码（常量、模型、方法）
- 更新项目文档
- 更新版本记录
- 准备 git 提交

**适用场景：**
- 添加微信支付新接口
- 添加支付宝新接口
- 添加其他支付平台接口

## 安装 Skills

Skills 需要安装到全局 `~/.claude/skills/` 目录才能被 Claude Code 识别。

### 方法一：使用脚本安装（推荐）

```bash
# 在项目根目录执行
.claude/skills/install.sh
```

### 方法二：手动安装

```bash
# 复制 skill 到全局目录
cp -r .claude/skills/add-payment-interface ~/.claude/skills/

# 重新加载 skills
# 在 Claude Code 中执行: /skills reload
```

## 使用 Skills

安装完成后，在 Claude Code 中：

1. 查看所有可用 skills：
   ```
   /skills
   ```

2. 使用 skill：
   - 直接告诉 Claude："使用 add-payment-interface skill 添加 xxx 接口"
   - 或者提供接口文档链接，Claude 会自动识别并使用相应的 skill

## 维护说明

### 更新 Skill

如果需要修改 skill：

1. 编辑项目中的 skill 文件：`.claude/skills/add-payment-interface/SKILL.md`
2. 重新安装到全局目录：
   ```bash
   cp -r .claude/skills/add-payment-interface ~/.claude/skills/
   ```
3. 在 Claude Code 中重新加载：`/skills reload`

### 换电脑后的设置

1. Clone 项目仓库
2. 运行安装脚本：
   ```bash
   .claude/skills/install.sh
   ```
3. 在 Claude Code 中验证：`/skills`

## 目录结构

```
.claude/
└── skills/
    ├── README.md                          # 本文件
    ├── install.sh                         # 安装脚本
    └── add-payment-interface/
        └── SKILL.md                       # Skill 定义文件
```

## 注意事项

- Skills 必须安装到全局目录 `~/.claude/skills/` 才能生效
- 项目中的 `.claude/skills/` 目录仅用于版本控制和分发
- 修改 skill 后需要重新安装并重新加载
- 建议将 `.claude/skills/` 目录提交到 git，方便团队协作
