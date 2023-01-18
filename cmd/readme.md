
### `/cmd`

Principais aplicações para este projeto.

O nome do diretório para cada aplicação deve corresponder ao nome do executável que você deseja ter (ex. `/cmd/myapp`).

Não coloque muitos códigos no diretório da aplicação. Se você acha que o código pode ser importado e usado em outros projetos, ele deve estar no diretório `/pkg`. Se o código não for reutilizável ou se você não quiser que outros o reutilizem, coloque esse código no diretório `/internal`. Você ficará surpreso com o que os outros farão, então seja explícito sobre suas intenções!

É comum ter uma pequena função `main` que importa e invoca o código dos diretórios` /internal` e `/pkg` e nada mais.

https://github.com/golang-standards/project-layout