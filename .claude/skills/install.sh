#!/bin/bash

# GoPay Skills 安装脚本
# 将项目中的 skills 安装到全局 ~/.claude/skills/ 目录

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
GLOBAL_SKILLS_DIR="$HOME/.claude/skills"

echo "🚀 开始安装 GoPay Skills..."

# 创建全局 skills 目录（如果不存在）
if [ ! -d "$GLOBAL_SKILLS_DIR" ]; then
    echo "📁 创建全局 skills 目录: $GLOBAL_SKILLS_DIR"
    mkdir -p "$GLOBAL_SKILLS_DIR"
fi

# 安装每个 skill
for skill_dir in "$SCRIPT_DIR"/*/; do
    if [ -d "$skill_dir" ]; then
        skill_name=$(basename "$skill_dir")

        # 跳过非 skill 目录
        if [ ! -f "$skill_dir/SKILL.md" ]; then
            continue
        fi

        echo "📦 安装 skill: $skill_name"

        # 复制到全局目录
        cp -r "$skill_dir" "$GLOBAL_SKILLS_DIR/"

        echo "   ✅ $skill_name 安装完成"
    fi
done

echo ""
echo "✨ 所有 skills 安装完成！"
echo ""
echo "📝 下一步："
echo "   1. 在 Claude Code 中执行: /skills reload"
echo "   2. 执行 /skills 查看已安装的 skills"
echo ""
