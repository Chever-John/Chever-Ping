# 关于 commit 规范化 CI 工具的选型思考

[![jaywcjlove/sb](https://jaywcjlove.github.io/sb/lang/chinese.svg)](project_design_methodology.md)

[English](../../en-US/thinkings/project_design_methodology.md) | [中文](project_design_methodology.md)

一个好的 CI 工具，能够在一定程度上规范开发者们的行为。基于此目标，本文章将展开对 CI 工具的选型思考。

## 思考角度

从一些已经常用的工具，到目前最火的 openAI 的工具，都尝试一遍。

- 已经很常用的工具
  - [commitizen-go](https://github.com/lintingzhen/commitizen-go)：使你进入交互模式，并根据提示生成 Commit Message，然后提交；
  - commit-msg：githooks，在 commit-msg 中，指定检查的规则，commit-msg 是个脚本，可以根据需要自己写脚本实现；
  - [go-gitlint](https://github.com/llorllale/go-gitlint)：检查历史提交的 Commit Message 是否符合 Angular 规范，可以将该工具添加在 CI 流程中，确保 Commit Message 都是符合规范的；
  - [gsemver](https://github.com/arnaud-deprez/gsemver)：语义化版本自动生成工具。
- 基于 ChatGPT 的工具
  - [commitgpt](https://github.com/RomanHotsiy/commitgpt)：有 1.1k 的 stars；
  - [gpt-commit](https://github.com/markuswt/gpt-commit)：python 写的工具，有 225 stars。
