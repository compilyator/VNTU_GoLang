# Git із командного рядка: команди, стани та типові сценарії

Ця пам'ятка пояснює не лише синтаксис команд Git, а й умови, за яких їх потрібно виконувати. Команди наведено для звичайного навчального репозиторію з основною гілкою `main` і віддаленим репозиторієм `origin`.

Git-команди однакові в PowerShell, Git Bash і Bash. Відрізняються переважно команди оболонки для переходу між каталогами та роботи з файлами.

> **Головне правило:** перед командою, що змінює індекс, історію або віддалений репозиторій, виконайте `git status`, перевірте поточний каталог і зрозумійте, які саме файли та коміти будуть змінені.

## 1. Модель Git

```text
робоче дерево --git add--> staging area --git commit--> локальна історія
      ^                                                    |
      |                                                    | git push
      |                                                    v
      +---------------- git pull / merge -------- віддалений репозиторій
```

- **Робоче дерево (working tree)** — файли, які ви зараз редагуєте.
- **Індекс, або staging area** — підготовлений знімок майбутнього коміту.
- **Коміт** — зафіксований знімок підготовлених змін із автором, часом, повідомленням і посиланням на попередній коміт.
- **Гілка** — рухоме ім'я, що вказує на один із комітів.
- **`HEAD`** — посилання на поточну гілку або безпосередньо на поточний коміт.
- **Remote** — збережена адреса іншого репозиторію. Назва `origin` є домовленістю, а не спеціальним сервером.

`git add` не надсилає файли на GitHub, а `git commit` не виконує `push`. Це три різні операції.

## 2. Заповнювачі в командах

У записах нижче конструкції `<...>` є заповнювачами:

```text
<file>        конкретний файл, наприклад main.go
<branch>      назва гілки, наприклад feature/validation
<commit>      хеш або інше посилання на коміт
<remote-url>  повна HTTPS- або SSH-адреса репозиторію
```

Кутові дужки не вводяться. Наприклад:

```powershell
git add -- main.go
```

а не:

```powershell
git add -- <main.go>
```

Подвійний дефіс `--` завершує параметри Git. Після нього значення трактується як шлях, навіть якщо ім'я починається з дефіса.

## 3. Перевірка встановлення і довідка

```powershell
git --version
git help <command>
git <command> --help
git <command> -h
```

- `git --version` показує встановлену версію.
- `git help commit` відкриває повну документацію команди.
- `git commit -h` показує коротку довідку в консолі.

Не копіюйте команду з LLM або вебсторінки, якщо не можете пояснити її параметри й цільові файли.

## 4. Початкове налаштування автора

Переглянути поточні значення:

```powershell
git config --global --get user.name
git config --global --get user.email
git config --global --list
```

Установити ім'я й email для всіх локальних репозиторіїв користувача:

```powershell
git config --global user.name "Student Name"
git config --global user.email "student@example.com"
```

Установити значення тільки для поточного репозиторію:

```powershell
git config user.name "Student Name"
git config user.email "student@example.com"
```

- `--global` змінює конфігурацію поточного користувача.
- Без `--global` команда діє лише в поточному репозиторії.
- Email потрапляє в метадані комітів. Для публічного репозиторію можна використати GitHub-provided `noreply` email, налаштований у GitHub.

## 5. Створення локального репозиторію

Спочатку перейдіть саме до кореня проєкту:

```powershell
cd <project-directory>
pwd
ls
git init -b main
git status
```

- `init` створює службовий каталог `.git` у поточному каталозі.
- `-b main` одразу задає назву початкової гілки.
- `git init` не потрібно повторювати для кожної лабораторної або версії програми.

Якщо встановлена версія Git не підтримує `-b`:

```powershell
git init
git branch -M main
```

- `branch -M main` перейменовує поточну гілку на `main`.
- `-M` дозволяє примусове перейменування; тому спочатку перевірте поточну гілку.

Не виконуйте `git init` у домашньому каталозі, батьківському каталозі з кількома проєктами або на всьому навчальному диску.

## 6. Щоденний безпечний цикл

```powershell
git status
git diff
git add -- main.go analyzer/analyzer.go
git diff --staged
git commit -m "feat(cli): validate analyze arguments"
git status
git push
```

### `git status`

Показує:

- поточну гілку;
- staged changes;
- unstaged changes;
- невідстежувані файли;
- розходження з upstream-гілкою.

Корисна компактна форма:

```powershell
git status --short --branch
```

- `--short` використовує компактні двосимвольні стани.
- `--branch` додає назву гілки та інформацію про upstream.

### `git diff`

```powershell
git diff
git diff -- main.go
git diff --staged
git diff --staged -- analyzer/analyzer.go
```

- Без параметрів показує незастейджені зміни tracked files.
- `-- <file>` обмежує результат конкретним шляхом.
- `--staged`, або `--cached`, показує майбутній коміт.
- Невідстежуваний файл не з'явиться у звичайному `git diff`, доки його не додано до індексу.

### `git add`

```powershell
git add -- main.go
git add -- main.go go.mod go.sum
git add -p -- analyzer/analyzer.go
```

- `git add -- <paths>` додає поточний вміст названих шляхів до staging area.
- `-p`, або `--patch`, дозволяє по черзі вибирати окремі фрагменти змін.
- Повторний `git add` потрібний після нового редагування вже застейдженого файла.

Для навчальної роботи віддавайте перевагу точним шляхам. `git add .` і `git add -A` можуть непомітно додати звіти, архіви, секрети, виконувані файли або налаштування IDE.

### `git commit`

```powershell
git commit -m "fix(analyzer): reject values outside variant range"
```

- `-m` задає коротке повідомлення без відкриття редактора.
- Коміт містить лише staged snapshot.
- Unstaged та untracked files до нього не потрапляють.

Рекомендований формат повідомлень:

```text
<type>[optional scope]: <imperative description>
```

Приклади:

```text
feat(cli): add analyze command
fix(storage): preserve file after failed write
test(analyzer): cover category boundaries
docs(readme): explain local setup
refactor(model): extract domain types
```

Один коміт має представляти одну логічно завершену зміну. Він може містити кілька файлів, якщо вони потрібні для одного результату.

## 7. Перегляд історії

```powershell
git log
git log --oneline
git log --graph --oneline --all --decorate
git show <commit>
git show --stat <commit>
git diff <older-commit>..<newer-commit>
```

- `--oneline` показує скорочений хеш і заголовок.
- `--graph` малює структуру гілок.
- `--all` включає всі локально відомі гілки.
- `--decorate` показує назви гілок і тегів біля комітів.
- `show` показує коміт і його patch.
- `--stat` замість повного patch показує список і обсяг змінених файлів.

## 8. Віддалені репозиторії

```powershell
git remote -v
git remote add origin <remote-url>
git remote get-url origin
git remote set-url origin <new-url>
git fetch origin
git push -u origin main
git push
```

- `remote -v` показує адреси для fetch і push.
- `remote add origin URL` додає remote з назвою `origin`.
- `remote get-url origin` показує одну фактичну адресу.
- `remote set-url` змінює помилкову або застарілу адресу.
- `fetch` отримує коміти й remote-tracking branches, але не змінює робочі файли.
- `push -u origin main` надсилає `main` і встановлює upstream.
- Після `-u` зазвичай достатньо `git push` і `git pull`.

Якщо `origin` уже існує, не додавайте другий remote з тією самою назвою. Спочатку виконайте `git remote -v`, а за потреби — `git remote set-url`.

## 9. `fetch`, `pull` і `push`

### Безпечна перевірка чужих змін

```powershell
git fetch origin
git log --oneline --graph --decorate HEAD..origin/main
git diff HEAD..origin/main
```

Після перегляду можна включити зміни:

```powershell
git merge origin/main
```

### Звичайний `pull`

```powershell
git pull --ff-only
```

- `pull` виконує fetch і потім інтеграцію.
- `--ff-only` дозволяє лише fast-forward і відмовляється створювати неочікуваний merge commit.
- Виконуйте команду в чистому робочому дереві після `git status`.

Якщо команда відмовила через локальні й віддалені незалежні коміти, спочатку перегляньте історію. Не додавайте навмання `--force` або `--allow-unrelated-histories`.

### Надсилання

```powershell
git push
```

Push може бути відхилено, якщо remote містить нові коміти. Тоді:

```powershell
git status
git fetch origin
git log --graph --oneline --all --decorate
```

Після усвідомленої інтеграції remote changes повторіть push.

## 10. Клонування

```powershell
cd <parent-directory>
git clone <remote-url>
git clone <remote-url> <local-directory-name>
```

- `clone` створює новий каталог, отримує історію, налаштовує `origin` і checkout основної гілки.
- Другий аргумент необов'язково задає локальну назву каталогу.
- Не виконуйте clone всередину іншого репозиторію без усвідомленої потреби.

Після клонування:

```powershell
cd <local-directory-name>
git status
git remote -v
git log --oneline -5
```

## 11. Гілки

```powershell
git branch
git switch -c feature/input-validation
git switch main
git merge feature/input-validation
git branch -d feature/input-validation
```

- `branch` показує локальні гілки; `*` позначає поточну.
- `switch -c` створює гілку від поточного коміту й переходить до неї.
- `switch` без `-c` переходить до наявної гілки.
- `merge` включає зміни в поточну гілку, тому перед ним перевірте, де перебуваєте.
- `branch -d` видаляє вже об'єднану локальну гілку; `-D` є примусовим і може втратити доступ до незлитих комітів.

Типовий сценарій:

```powershell
git switch main
git pull --ff-only
git switch -c feature/input-validation
# редагування, тести, add і commit
git switch main
git merge feature/input-validation
git push
```

## 12. Конфлікт злиття

Конфлікт означає, що Git не може самостійно вибрати правильне поєднання змін.

```powershell
git status
```

Відкрийте конфліктні файли й знайдіть маркери. Нижче символ `│` додано лише для безпечного показу; у справжньому конфліктному файлі рядки починатимуться без нього:

```text
│ <<<<<<< HEAD
поточна версія
│ =======
інша версія
│ >>>>>>> branch-name
```

Вручну сформуйте правильний підсумковий текст і видаліть усі маркери. Потім:

```powershell
git add -- <resolved-file>
git status
git diff --staged
git commit
```

Якщо конфлікт виник під час rebase, Git підкаже `git rebase --continue`. Не використовуйте цю команду під час звичайного merge.

Щоб скасувати незавершене злиття до коміту:

```powershell
git merge --abort
```

Перед `--abort` прочитайте `git status` і переконайтеся, що не втратите потрібні незбережені зміни.

## 13. Виправлення помилок без переписування спільної історії

### Прибрати файл зі staging area, зберігши його зміни

```powershell
git restore --staged -- <file>
```

### Відкинути незастейджені зміни файла

```powershell
git restore -- <file>
```

Ця команда перезаписує робочий файл версією з індексу. Незакомічені зміни можуть бути втрачені — перед виконанням перегляньте `git diff -- <file>`.

### Скасувати вже опублікований коміт новим комітом

```powershell
git revert <commit>
```

`revert` створює новий коміт із протилежними змінами й зберігає історію. Це типовий безпечний спосіб скасування опублікованого коміту.

Не використовуйте `git reset --hard`, `git clean -fd` або force push як універсальне «виправлення». Вони можуть незворотно видалити локальні дані або переписати історію інших людей.

## 14. `.gitignore`

Файл `.gitignore` описує невідстежувані файли, які Git не повинен пропонувати до коміту.

Приклад для невеликого Go-проєкту:

```gitignore
# Виконувані файли й результати збирання
*.exe
*.test
*.out
dist/
bin/

# Секрети та локальна конфігурація
.env
.env.*
!.env.example
*.pem
*.key

# IDE та ОС
.vscode/
.idea/
.DS_Store
Thumbs.db

# Звіти й архіви з персональними даними
reports/
*.zip
```

Перевірити правило:

```powershell
git check-ignore -v -- <file>
```

Важливо: `.gitignore` не видаляє файл, який уже відстежується. Щоб припинити його відстеження, але залишити локально:

```powershell
git rm --cached -- <file>
```

Після цього створіть коміт. Якщо файл містив секрет, цього недостатньо: секрет уже може бути в історії та чужих клонах, тому його потрібно негайно відкликати або замінити.

## 15. Перевірка перед комітом і push

```powershell
git status --short --branch
git diff
git diff --staged
git diff --staged --name-only
git remote -v
```

Перевірте:

- чи ви в потрібному каталозі та гілці;
- чи немає `.env`, приватних ключів, токенів, паролів або реальних персональних даних;
- чи не додано звіти, скриншоти, архіви, виконувані файли й конфігурацію IDE;
- чи staged diff відповідає одному повідомленню коміту;
- чи `origin` вказує на ваш репозиторій;
- чи пройшли форматування й тести.

Не вставляйте секрет у командний рядок, повідомлення коміту, назву гілки або URL remote. Історія команд оболонки також може зберігатися.

## 16. Швидкий вибір команди за ситуацією

| Ситуація | Спочатку | Основна команда |
| --- | --- | --- |
| Перевірити стан | Перейти до кореня проєкту | `git status --short --branch` |
| Побачити незастейджені зміни | Перевірити status | `git diff` |
| Підготувати конкретний файл | Переглянути diff | `git add -- <file>` |
| Перевірити майбутній коміт | Додати потрібні файли | `git diff --staged` |
| Створити коміт | Перевірити staged diff і тести | `git commit -m "..."` |
| Переглянути історію | — | `git log --graph --oneline --all --decorate` |
| Отримати remote changes без зміни файлів | Перевірити remote | `git fetch origin` |
| Оновитися лише fast-forward | Чисте дерево, правильна гілка | `git pull --ff-only` |
| Перший push гілки | Є локальні коміти | `git push -u origin <branch>` |
| Наступний push | Upstream уже задано | `git push` |
| Прибрати файл зі staging | Перевірити staged diff | `git restore --staged -- <file>` |
| Скасувати опублікований коміт | Перевірити коміт | `git revert <commit>` |
| Перевірити правило ignore | Файл ще не додано | `git check-ignore -v -- <file>` |

## 17. Джерела

Перевірено 19.09.2026:

- [офіційна документація Git](https://git-scm.com/docs);
- [офіційний Git Cheat Sheet](https://git-scm.com/cheat-sheet.pdf);
- [довідка GitHub про ігнорування файлів](https://docs.github.com/en/get-started/getting-started-with-git/ignoring-files);
- [довідка GitHub про підключення локального коду](https://docs.github.com/en/migrations/importing-source-code/using-the-command-line-to-import-source-code/adding-locally-hosted-code-to-github).

[До переліку практичних інструкцій](../README.md) · [Реєстрація та публічні репозиторії GitHub](../github/README.md)
