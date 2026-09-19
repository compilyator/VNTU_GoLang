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

## 5. Підготовка локального проєкту до публікації

Перед `git add` або `push`:

```powershell
cd <project-directory>
pwd
ls
git status --short --branch
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

Після pull перевірте проєкт і створіть перший змістовний коміт:

```powershell
git status
git diff
git add -- .gitignore main.go go.mod go.sum
git diff --staged
git commit -m "feat(analyzer): add initial console analyzer"
git log --graph --oneline --all --decorate
git push -u origin main
```

Якщо `.gitignore` уже прийшов із GitHub і не змінювався, він не з'явиться в новому staged snapshot — це нормально.

## 7. Сценарій B: локальний Git-репозиторій уже має коміти

Офіційна документація GitHub радить для такого імпорту створити **порожній** remote: не додавати README, `.gitignore` або license на сторінці створення. Це запобігає появі двох незалежних історій.

На GitHub створіть public repository без початкових файлів. Потім:

```powershell
cd <project-directory>
git status --short --branch
git log --oneline -5
git remote add origin https://github.com/<username>/<repository>.git
git remote -v
git push -u origin main
```

Якщо локальна гілка має іншу назву:

```powershell
git branch -M main
git push -u origin main
```

Не виконуйте `git pull` із порожнього remote: там немає гілки або комітів для отримання.

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

Не виконуйте `git init` після звичайного clone: репозиторій уже ініціалізовано.

## 9. Сценарій D: підключення remote, який уже має власні коміти

Спочатку не робіть pull навмання:

```powershell
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

## 10. Якщо `origin` налаштовано неправильно

```powershell
git remote -v
git remote get-url origin
git remote set-url origin https://github.com/<username>/<correct-repository>.git
git remote -v
```

Не надсилайте код, доки owner і repository не перевірені. Помилковий push у чужий або публічний репозиторій може розкрити дані.

## 11. Перевірка перед першим і кожним наступним push

```powershell
git status --short --branch
git diff
git diff --staged
git diff --staged --name-only
git log --oneline -5
git remote -v
```

Контрольний список:

- поточний каталог є коренем потрібного проєкту;
- гілка й remote правильні;
- `.gitignore` існує та перевірений;
- staged files переглянуті поіменно;
- у diff немає паролів, токенів, ключів, cookies, приватних URL і персональних даних;
- у commit message немає секретів;
- тести й форматування пройшли;
- remote URL не містить token;
- repository visibility справді `Public` і ви готові показати весь staged content будь-кому.

GitHub push protection і secret scanning є додатковими бар'єрами, але не замінюють ручну перевірку. Вони не гарантують виявлення кожного формату секрету.

## 12. Що робити після випадкової публікації секрету

1. **Негайно відкличте або замініть секрет** у сервісі, який його видав.
2. Не покладайтеся на видалення файла останнім комітом: попередній commit усе ще містить значення.
3. Не публікуйте секрет повторно в issue, повідомленні коміту, чаті чи скриншоті під час прохання про допомогу.
4. Повідомте викладача або відповідального адміністратора без передавання самого секрету.
5. Лише після відкликання оцініть потребу очищення історії за офіційною процедурою GitHub.
6. Перевірте forks, clones, CI logs, releases та інші місця, куди секрет міг потрапити.

Переписування історії складне: воно змінює commit hashes, потребує координації з усіма клонами й не відкликає вже скопійовані credentials. Офіційна інструкція: [Removing sensitive data from a repository](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/removing-sensitive-data-from-a-repository).

## 13. Типові помилки

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

## 14. Джерела

Перевірено 19.09.2026:

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
