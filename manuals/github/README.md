# GitHub: реєстрація, публічний репозиторій, clone і підключення локального проєкту

Ця інструкція охоплює створення облікового запису GitHub, безпечну публікацію навчального проєкту, клонування готового репозиторію та підключення GitHub до каталогу, який уже містить код.

> **Беззаперечне правило:** усе, що потрапило в інтернет, потрібно вважати таким, що залишилося там назавжди. Видалення файла, коміту або репозиторію не гарантує видалення копій, кешів, форків, клонів, логів чи вже скопійованих секретів. Якщо секрет було опубліковано, його спочатку потрібно відкликати або замінити, а не лише видалити з останньої версії файла.

## 1. Що означає `Public`

Публічний репозиторій можуть переглядати, клонувати й аналізувати сторонні люди та автоматизовані системи. Публічність стосується не лише поточних файлів, а й доступної історії комітів, гілок, тегів, issues, pull requests і прикріплених матеріалів.

Не публікуйте:

- паролі, API keys, access tokens, session cookies і recovery codes;
- приватні SSH keys, TLS keys, seed phrases або файли credentials;
- `.env` із реальними значеннями;
- звіти з ПІБ, номером групи, адресою, телефоном або приватними шляхами;
- дампи баз даних, журнали з токенами, production-конфігурацію;
- файли, ліцензія яких не дозволяє публікацію;
- відповіді, службові матеріали викладача чи інші дані, не призначені для студентського репозиторію.

Навіть private repository не є сховищем секретів. Секрети передають через environment variables або спеціалізовані secret stores, а в репозиторії залишають лише безпечні приклади на кшталт `.env.example`.

## 2. Створення облікового запису

1. Відкрийте [форму реєстрації GitHub](https://github.com/signup).
2. Створіть особистий обліковий запис і виберіть унікальний username.
3. Використайте унікальний сильний пароль або підтримуваний social login.
4. Підтвердьте email: без підтвердженої адреси частина базових операцій, зокрема створення репозиторію, може бути недоступною.
5. Увімкніть two-factor authentication (2FA) і безпечно збережіть recovery codes поза репозиторіями та навчальними звітами.
6. За можливості додайте passkey або security key.
7. Перегляньте публічні поля профілю й не заповнюйте дані, які не хочете розкривати.

Для комітів можна використовувати GitHub-provided `noreply` email. Налаштування email коміту перевіряють у GitHub `Settings` → `Emails`, а локальне значення — через:

```powershell
git config --global --get user.email
```

## 3. Автентифікація Git у командному рядку

GitHub не приймає пароль облікового запису як пароль для Git over HTTPS. Для початківця рекомендований HTTPS із Git Credential Manager або GitHub CLI, які відкривають безпечний browser login і зберігають credentials через системні механізми.

### HTTPS через Git Credential Manager

Після встановлення сучасного Git for Windows Git Credential Manager зазвичай уже доступний. Під час першого `clone`, `pull` або `push` підтвердьте вхід у браузері.

Не вставляйте personal access token у remote URL:

```text
https://TOKEN@github.com/user/repository.git   # небезпечно
```

Такий URL може потрапити в `.git/config`, історію оболонки, логи чи скриншоти.

### GitHub CLI

Якщо встановлено `gh`:

```powershell
gh auth login
gh auth status
```

Під час входу можна вибрати `GitHub.com`, `HTTPS` і browser authentication.

### SSH

SSH є повноцінною альтернативою HTTPS, але потребує окремої пари ключів і додавання **публічного** ключа до GitHub. Приватний ключ ніколи не завантажують у репозиторій і не надсилають іншій людині. Для першої навчальної роботи HTTPS із credential manager зазвичай простіший.

Офіційний огляд способів: [About authentication to GitHub](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/about-authentication-to-github).

## 4. Назва навчального репозиторію

Назва має описувати проєкт, а не номер лабораторної:

```text
weather-data-analyzer
air-quality-cli
book-stats
launch-intervals
```

Використовуйте lowercase і дефіси. Не включайте ПІБ, email, номер телефону, пароль, номер студентського квитка або інші персональні дані. Якщо один проєкт розвивається в кількох лабораторних, не створюйте `lab-02`, `lab-03`, `lab-04` як окремі репозиторії — продовжуйте одну історію.

## 5. Виберіть один сценарій

Це альтернативні способи початку роботи. Виконайте лише свій сценарій із розділів 6–9, а потім переходьте до спільної підготовки в розділі 10.

| Що вже є | Сценарій |
| --- | --- |
| Локальний каталог із кодом без комітів; потрібно почати з README на GitHub | A — розділ 6: `init` і отримання початкового коміту |
| Локальний Git-репозиторій із комітами; GitHub-репозиторію ще немає | B — розділ 7: порожній GitHub-репозиторій і підключення `origin` |
| Репозиторій на GitHub; локальної копії ще немає | C — розділ 8: `clone` |
| Локальний репозиторій із комітами та GitHub-репозиторій із власними комітами | D — розділ 9: перевірка сумісності історій |

Якщо починаєте з нуля, створіть репозиторій на GitHub за кроками «На GitHub» у сценарії A, а потім клонуйте його за сценарієм C і створюйте код у клонованому каталозі.

Команди нижче призначені для **PowerShell у Windows** зі встановленим Git. Позначення `<project-directory>`, `<username>` та `<repository>` замініть власними значеннями; кутові дужки не вводьте. Налаштування Git і основні поняття описано в [інструкції з Git](../git/README.md).

Перед `init` переконайтеся через `pwd` і `ls`, що відкрили каталог лише потрібного проєкту. Для `clone` відкрийте батьківський каталог: Git створить у ньому окрему папку.

## 6. Сценарій A: є локальний каталог із кодом, але ще немає локальних комітів

Цей сценарій зручний для навчальної роботи, якщо потрібно спочатку створити GitHub repository з README, отримати його перший коміт, а потім окремо зафіксувати готову локальну версію програми.

### На GitHub

1. У правому верхньому куті виберіть `+` → `New repository`.
2. Виберіть власний account як owner.
3. Уведіть тематичну назву.
4. Виберіть `Public`.
5. Увімкніть `Add a README file`.
6. За потреби виберіть Go `.gitignore`.
7. Не додавайте ліцензію, якщо не розумієте її умови або не отримали вказівку щодо вибору.
8. Натисніть `Create repository`.

### У локальному каталозі

```powershell
cd <project-directory>
pwd
ls
git init -b main
git remote add origin https://github.com/<username>/<repository>.git
git remote -v
git pull origin main
git status
```

- `git init -b main` створює локальний репозиторій із гілкою `main`.
- `remote add origin` зберігає адресу GitHub.
- `remote -v` дає змогу виявити помилковий owner або repository до push.
- `pull origin main` отримує початковий README. Локальна історія до цього повинна бути порожньою.

Якщо локально вже був `README.md` з іншими даними, pull може зупинитися, щоб не перезаписати untracked file. Не видаляйте файл навмання: порівняйте версії, перейменуйте або узгодьте їх, а потім повторіть дію.

Після pull перейдіть до розділу 10: перевірте файли, налаштуйте `.gitignore` і створіть коміт із кодом.

## 7. Сценарій B: локальний Git-репозиторій уже має коміти

Офіційна документація GitHub радить для такого імпорту створити **порожній** remote: не додавати README, `.gitignore` або license на сторінці створення. Це запобігає появі двох незалежних історій.

На GitHub створіть public repository без початкових файлів. Потім:

```powershell
cd <project-directory>
git status --short --branch
git log --oneline -5
git remote add origin https://github.com/<username>/<repository>.git
git remote -v
```

Якщо локальна гілка має іншу назву:

```powershell
git branch -M main
```

Не виконуйте `git pull` із порожнього remote: там немає гілки або комітів для отримання. Перед публікацією перейдіть до розділу 10 і перевірте також наявну історію: push передає попередні коміти, а не лише поточні файли.

## 8. Сценарій C: GitHub-репозиторій уже існує, локальної копії немає

Скопіюйте HTTPS URL на сторінці репозиторію (`Code` → `HTTPS`) і виконайте:

```powershell
cd <parent-directory>
git clone https://github.com/<username>/<repository>.git
cd <repository>
git status
git remote -v
git log --oneline -5
```

`clone` одночасно:

- створює локальний каталог;
- завантажує доступну історію;
- створює remote `origin`;
- налаштовує tracking основної гілки;
- виконує checkout робочих файлів.

Не виконуйте `git init` після звичайного clone: репозиторій уже ініціалізовано. Якщо потрібно лише отримати копію, роботу завершено. Якщо будете додавати власний код, працюйте в клонованому каталозі й перейдіть до розділу 10.

## 9. Сценарій D: підключення remote, який уже має власні коміти

Відкрийте корінь локального проєкту. Спочатку не робіть pull навмання:

```powershell
cd <project-directory>
git status --short --branch
git remote add origin https://github.com/<username>/<repository>.git
git remote -v
git fetch origin
git log --graph --oneline --all --decorate
```

Якщо локальна й remote-гілки мають спільного предка, інтегруйте зміни відповідно до правил проєкту, наприклад:

```powershell
git pull --rebase origin main
```

Якщо історії незалежні, зупиніться й визначте, який сценарій мав бути використаний. `--allow-unrelated-histories`, force push або видалення `.git` не є стандартним виправленням. Для навчального проєкту часто безпечніше створити правильний порожній remote або повторити процедуру з резервної копії після консультації викладача.

Після узгодження історій перейдіть до розділу 10. Якщо origin уже існує, перевірте його за розділом 11 замість повторного emote add.

## 10. Підготовка файлів і створення коміту

Після обраного сценарію локальний Git-репозиторій уже має існувати. Перед першим додаванням власних файлів перевірте його:

```powershell
cd <project-directory>
pwd
ls
git status --short --branch
git rev-parse --show-toplevel
```

Перевірте, що корінь репозиторію міститиме тільки проєкт, а не домашній каталог чи каталог з іншими роботами.

### Мінімальний `.gitignore` для Go

Створіть `.gitignore` до першого `git add`:

```gitignore
# Виконувані файли та результати тестів
*.exe
*.exe~
*.test
*.out
coverage.*
bin/
dist/

# Секрети й локальна конфігурація
.env
.env.*
!.env.example
*.pem
*.key
credentials.json

# IDE та ОС
.vscode/
.idea/
.DS_Store
Thumbs.db

# Звіти й архіви, які не належать до коду
reports/
*.zip
```

`.gitignore` потрібно адаптувати до проєкту. Не ігноруйте `go.mod`, `go.sum`, вихідний код або корисні тестові fixtures лише для того, щоб приховати проблему.

Перевірка конкретного файла:

```powershell
git check-ignore -v -- .env
```

Перегляд уже відстежуваних файлів:

```powershell
git ls-files
```

`.gitignore` не діє заднім числом на tracked files.

### Перевірте код і створіть коміт

Для Go-проєкту виконайте форматування та тести з кореня модуля:

```powershell
go fmt ./...
go test ./...
```

Потім перегляньте зміни та додайте лише потрібні файли:

```powershell
git status
git diff
git add -- .gitignore main.go go.mod
git diff --staged
git diff --staged --name-only
git commit -m "feat(analyzer): add initial console analyzer"
git log --graph --oneline --all --decorate
```

Перелік у `git add` є прикладом: замініть його фактичними файлами проєкту. Додайте `go.sum`, якщо він існує, а також потрібні пакети й тести. Якщо `.gitignore` отримано з GitHub і не змінено, він не потрапить до нового коміту — це нормально. Якщо власних змін немає, новий коміт не потрібний.

Для сценаріїв B і D перегляньте також попередні коміти, які буде передано на GitHub. Чистий робочий каталог і новий `.gitignore` не гарантують відсутності секретів в історії.

Якщо використовуєте LLM, онлайн-діалог доречний для пояснення команд і аналізу повідомлень про помилки. Передавайте лише знеособлені фрагменти без секретів; перевіряйте запропоновані команди за документацією та самостійно переглядайте `git diff`. Для навчального звіту запишіть сервіс, модель (якщо відома), дату, мету запиту, використаний результат, власні виправлення і спосіб перевірки.

Після створення коміту перевірте `origin` у розділі 11, а потім виконайте перевірку та публікацію за розділом 12.

## 11. Якщо `origin` налаштовано неправильно

```powershell
git remote -v
git remote get-url origin
git remote set-url origin https://github.com/<username>/<correct-repository>.git
git remote -v
```

Не надсилайте код, доки owner і repository не перевірені. Помилковий push у чужий або публічний репозиторій може розкрити дані.

## 12. Перевірка та публікація на GitHub

```powershell
git status --short --branch
git diff
git diff --staged
git diff --staged --name-only
git log --oneline -5
git remote -v
```

Після commit staged diff зазвичай порожній. Для перегляду останнього коміту використайте `git show --stat` та `git show`; якщо надсилаєте кілька комітів, перегляньте кожен із них, а не лише останній.

Контрольний список:

- поточний каталог є коренем потрібного проєкту;
- гілка й remote правильні;
- `.gitignore` існує та перевірений;
- staged files переглянуті поіменно;
- перевірені зміни й коміти, які надсилаєте, не містять паролів, токенів, ключів, cookies, приватних URL і персональних даних;
- у commit message немає секретів;
- тести й форматування пройшли;
- remote URL не містить token;
- repository visibility справді `Public` і ви готові показати всі файли та історію, які надсилаєте, будь-кому.

GitHub push protection і secret scanning є додатковими бар'єрами, але не замінюють ручну перевірку. Вони не гарантують виявлення кожного формату секрету.

Після перевірки надішліть коміти:

```powershell
git push -u origin main
```

Ця команда передбачає гілку `main`; перевірте її назву через `git branch --show-current`. Після першого успішного push із `-u` наступні публікації з цієї гілки можна виконувати через `git push`.

Відкрийте репозиторій на GitHub і перевірте, що в потрібній гілці з'явилися очікувані файли та останній коміт.

## 13. Що робити після випадкової публікації секрету

1. **Негайно відкличте або замініть секрет** у сервісі, який його видав.
2. Не покладайтеся на видалення файла останнім комітом: попередній commit усе ще містить значення.
3. Не публікуйте секрет повторно в issue, повідомленні коміту, чаті чи скриншоті під час прохання про допомогу.
4. Повідомте викладача або відповідального адміністратора без передавання самого секрету.
5. Лише після відкликання оцініть потребу очищення історії за офіційною процедурою GitHub.
6. Перевірте forks, clones, CI logs, releases та інші місця, куди секрет міг потрапити.

Переписування історії складне: воно змінює commit hashes, потребує координації з усіма клонами й не відкликає вже скопійовані credentials. Офіційна інструкція: [Removing sensitive data from a repository](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/removing-sensitive-data-from-a-repository).

## 14. Типові помилки

### `remote origin already exists`

Remote вже налаштований. Перевірте його:

```powershell
git remote -v
```

За потреби змініть URL через `git remote set-url`, а не додавайте другий `origin`.

### `src refspec main does not match any`

Зазвичай гілка `main` ще не має коміту або має іншу назву:

```powershell
git branch --show-current
git log --oneline -1
```

Створіть перевірений commit або узгодьте назву гілки; не створюйте порожні фіктивні коміти лише для приховування причини.

### `rejected ... non-fast-forward`

Remote має коміти, яких немає локально. Виконайте `git fetch origin`, перегляньте граф і свідомо інтегруйте зміни. Force push не є типовим розв'язанням.

### Git просить пароль у консолі

Пароль GitHub account не працює для Git over HTTPS. Використайте browser authentication через Git Credential Manager або `gh auth login`; personal access token не вставляйте в remote URL.

### `.gitignore` не приховує файл

Файл уже tracked. Перевірте `git ls-files`, видаліть його лише з індексу через `git rm --cached -- <file>` і створіть commit. Якщо там був секрет — спочатку revoke/rotate.

## 15. Джерела

Джерела загальних налаштувань перевірено 19.09.2026. Послідовність init → підготовка → commit → push та поведінку clone повторно звірено з офіційною документацією 06.10.2026:

- [команда git clone](https://git-scm.com/docs/git-clone) і [команда git init](https://git-scm.com/docs/git-init);
- [створення облікового запису GitHub](https://docs.github.com/en/account-and-profile/how-tos/account-management/creating-an-account-on-github);
- [початкове налаштування account, email і 2FA](https://docs.github.com/en/get-started/onboarding/getting-started-with-your-github-account);
- [створення нового репозиторію](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-new-repository);
- [додавання наявного локального коду до GitHub](https://docs.github.com/en/migrations/importing-source-code/using-the-command-line-to-import-source-code/adding-locally-hosted-code-to-github);
- [автентифікація в GitHub](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/about-authentication-to-github);
- [Git Credential Manager і GitHub CLI](https://docs.github.com/en/get-started/git-basics/caching-your-github-credentials-in-git);
- [ігнорування файлів](https://docs.github.com/en/get-started/getting-started-with-git/ignoring-files);
- [push protection](https://docs.github.com/en/code-security/concepts/secret-security/push-protection);
- [видалення чутливих даних з історії](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/removing-sensitive-data-from-a-repository).

[До переліку практичних інструкцій](../README.md) · [Git cheat sheet](../git/README.md)
