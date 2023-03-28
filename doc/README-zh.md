# Chever-Ping

[![jaywcjlove/sb](https://jaywcjlove.github.io/sb/lang/chinese.svg)](README-zh.md)

[English](../README.md) | [中文](README-zh.md)

## 项目背景

在这个项目中，我将主要编写一些运行在容器中的演示程序。我希望这个项目可以成为一个通用的适应容器开发的Web框架，以便我能够基于它快速实现任何我想要实践的想法。

然而，这个项目的核心是其优雅的实现方式。因此，我打算在这个项目上投入我的软件开发经验，并不断思考如何做到更优雅的实现。

## 项目路线图

- 规范设计
  - **项目文档规范**：包括用户文档和开发文档，并输出 doc 目录的框架结构。
  - **API 接口文档规范**
  - **版本规范**：遵守语义化版本规范或 SemVer，并发布[版本化思考文章](devel/zh-CN/thinkings/version_standardization.md)。
  - **commit 规范化**：撰写 [commit 自动化规范文章](devel/zh-CN/thinkings/version_standardization.md)，调研 OpenAI 的 commit CI 工具，选择 Angular 规范并最终实现 CI 规范化 commit。
    - 调研基于 OpenAI 的 commit CI 工具
    - 次而选择 Angular 规范
    - 最终实现 CI 规范化 commit
- **项目目录结构设计规范**
- 构建 Go 优雅项目设计方法论，输出一份[文章](devel/zh-CN/thinkings/project_design_methodology.md)覆盖以下的方法论。
- **构建 Makefile 管理项目**
- 完成 CI/CD
  - 流水线 pipeline 执行 make all 命令
  - 自动生成 CHANGELOG
  - 自动生成版本号

## 作者

- [@Chever-John](https://github.com/Chever-John)
