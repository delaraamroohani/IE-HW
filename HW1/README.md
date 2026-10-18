# تمرین اول
## سوال‌های مفهومی
###  ﺳﺆال ﯾﮏ: merge و rebase
با دستور  `git -h` به این تعاریف می‌رسیم:
```
merge      Join two or more development histories together
rebase     Reapply commits on top of another base tip
```

با دستور `git merge` آخرین نسخه یک برنچ را با برنچ خود ترکیب می‌کنیم. یک کامیت اضافه ایجاد می‌شود که تغییرات برنچ اصلی را روی برنچ خود اعمال میکند.
با دستور `git rebase` کامیت جدیدی اتفاق نمی‌افتد. مانند این است که گذشته را بازنویسی کنیم، به نحوی که بنظر بیاید برنچی که در حال حاضر روی آن هستیم ادامه آخرین نسجه برنچ اصلی است. 
[در Reddit یک نفر (jeenajeena) اینطور بیان کرده](https://www.reddit.com/r/git/comments/1qakug2/comment/nz497zn/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button):

>  I'm working on a project:
> 
> ```
> A-B-C
> ```
> 
> I spin off my branch for working on a feature, and I produce some work:
> ```
> A-B-C
>      \
>       1-2-3-4
> ```
> 
> So far, so good: my work is based on the most updated version of the project. I'm not missing any information.
> 
> The next day, some of my colleagues evolve the project with some important changes:
> ```
> A-B-C-D-E-F
>      \
>       1-2-3-4
> ```
> 
> Damn! D-E-F contain such important changes! If only I started working today, instead of yesterday, I would have of course spinned my feature branch off the version F, instead of on top of C. I mean, why on earth should I start on C, knowing this is an outdated version?
> 
> Well, if I rebase my branch on top of F it is like I started working right now:
> ```
> A-B-C-D-E-F
>           \
>            1-2-3-4
> ```
> It is really like restarting working from the scratch today. Only, it's done by Git, for free.
> 
> What I am going to push, conveys the information that, at the end of my activity, the entirety of my work 1-2-3-4 is a feature added on top of the last updated version of the project, that is, on F.
> 
> Wait! What if instead I merge F?
> ```
> A-B-C-D-E-F---
>      \        \
>       1-2-3-4-4F
> ```
> Amazing! The result 4F is exactly the same of 4 of the previous case. But this history conveys a very different information: it stresses on the fact that I started my work yesterday, on the (now outdated) version C of the project.

در صورتی که اهمیت داشته باشد که تیمی که با آن کار می‌کنم بدانند ابتدا با نسخه قدیمی‌تر برنچ اصلی کار می‌کردم، از `git merge` استفاده می‌کنم. در غیر این صورت می‌توانم از `git rebase` استفاده کنم تا تاریخچه پروژه تمیزتر باشد.
### سوال دو: انتخاب بخشی از تغییرات
بر اساس [مستندات GNU](https://www.gnu.org/software/diffutils/manual/html_node/Detailed-Unified.html)، هر بخش تغییرات یا hunk بخش حدالامکان مینیمال کد بین دو کامیت است که تغییر کرده. قبل از خطوط حذف شده - گذاشته می‌شود و قبل از خطوط اضافه شده + گذاشته می‌شود.
بر اساس [مستندات git](https://git-scm.com/book/en/v2/Git-Tools-Interactive-Staging)، برای interactive staging می‌توان از دستور `git add -i` استفاده کرد. در اینجا شبیه به `git status` اما دقیق‌تر وضعیت فایل‌ها را می‌توان دید و چند گزینه برای ما به نمایش گذاشته می‌شود، مانند زیر.
```console
$ git add -i
           staged     unstaged path
  1:    unchanged        +0/-1 TODO
  2:    unchanged        +1/-1 index.html
  3:    unchanged        +5/-1 lib/simplegit.rb

*** Commands ***
  1: [s]tatus     2: [u]pdate      3: [r]evert     4: [a]dd untracked
  5: [p]atch      6: [d]iff        7: [q]uit       8: [h]elp
What now>
```
با گزینه `p` برای patch و سپس انتخاب فایل یا فایل‌های مورد نظر، hunkها یک به یک برای ما به نمایش گذاشته می‌شوند و می‌توانیم هر کدام را یا برای staging انتخاب کنیم یا نکنیم. 
```console
Stage this hunk [y,n,a,d,/,j,J,g,e,?]? ?
y - stage this hunk
n - do not stage this hunk
a - stage this and all the remaining hunks in the file
d - do not stage this hunk nor any of the remaining hunks in the file
g - select a hunk to go to
/ - search for a hunk matching the given regex
j - leave this hunk undecided, see next undecided hunk
J - leave this hunk undecided, see next hunk
k - leave this hunk undecided, see previous undecided hunk
K - leave this hunk undecided, see previous hunk
s - split the current hunk into smaller hunks
e - manually edit the current hunk
? - print help
```
گاهی تغییرات زیادی انجام شده و برای انسجام کار و خوانایی history بهتر است این تغییرات به چند commit متفاوت تقسیم شوند. با استفاده از این ویژگی git می‌توان این خواسته را عملی کرد.
### سوال سوم: یافتن منشأ خطا
بر اساس [مستندات git](https://git-scm.com/docs/git-bisect)، برای کار با دستور `git bisect` ابتدا یک کامیت good و یک کامیت bad را مشخص می‌کنیم، و این دتسور به کمک جستجوی binary از ما می‌خواهد که کامیت‌های میانی را good و bad نامگذاری کنیم (نسخه checked out را کامپایل می‌کنیم و نتیجه را گزارش می‌دهیم) تا جایی که کامیت باگ‌ساز پیدا شود. دلیل کارایی این روش به سرچ باینری بودن آن است. به عنوان مثال ممکن است با ۶۷۵ کامیت طرف باشیم، اما پیدا کردن کامیت باگ‌ساز حدود ۱۰ مرحله طول خواهد کشید.
```
Bisecting: 675 revisions left to test after this (roughly 10 steps)
```
### سوال چهارم: بررسی تغییرات پیش از کامیت
1. وضعیت مخزن با `git status` بررسی می‌شود. برای دیدن تغییرات staged و unstaged از دستور `git add -i` استفاده می‌شود. تغییرات staged یعنی آماده کامیت هستند و تغییرات unstaged یا untracked تغییراتی هستند که قصد نداریم آنها را کامیت کنیم و اهمیتی به آنها نمی‌دهیم.
2. تغییرات وارد شده به بخش staged کامیت خواهند شد. 
3. فایلی به نام .gitignore داریم که می‌توانیم در آن فایل‌هایی که تحت version control نیستند را نگهداری کنیم. git به این فایل‌ها توجهی نمی‌کند و آنها هرگز کامیت نمی‌شوند. حائز اهمیت است که خود .gitignore مشخصا کامیت می‌شود. اگر فایلی از قبل tracked شده باشد گذاشتن آن در .gitignore تاثیری نخواهد داشت.
### سوال پنجم
1. دلیل conflict این است که جفت ما یک خط یکسان را به دو مقدار متفاوت تغییر داده‌ایم، و گیت نمی‌داند کدام تغییر را تغییر نهایی در نظر بگیرد. لزوما هر تغییری باعث conflict نمی‌شود، تنها تغییراتی که روی یک قطعه یکسان انجام شده باشند باعث conflict می‌شوند.
2. برای مشاهده conflictها از `git status` استفاده می‌شود. از فرمت زیر استفاده می‌کنند (برداشته از [مستندات گیت](https://git-scm.com/docs/git-merge)):
```
Here are lines that are either unchanged from the common
ancestor, or cleanly resolved because only one side changed,
or cleanly resolved because both sides changed the same way.
<<<<<<< yours:sample.txt
Conflict resolution is hard;
let's go shopping.
=======
Git makes conflict resolution easy.
>>>>>>> theirs:sample.txt
And here is another line that is cleanly resolved or unmodified.
```
بخشی که بین `<<<<<<<` و `=======` است بخش conflictی فایل ماست و بخشی که بین `=======` و `>>>>>>>` است مربوط به برنچی است که برنامه داریم به آن مرج کنیم. به اینها conflict marker می گویند. اطلاعات قبل و بین برای کانتکست است.
3. با `git status` فایل را پیدا می‌کنم و conflict markerها را بر میدارم. متن فایل را طوری که نیاز دارم تنظیم می‌کنم (مثلا جفت تغییرات را نگه میدارم). بعد دوباره از `git diff` و `git status` استفاده می‌کنم. اگر فایل مورد نظر در اینجا نمایش داده نشد، آن را دوباره `git add` می‌کنم و در نهایت `git commit`.
4. با اضافه کردن تگ `--abort`.
(در بخش ۳ و ۴ این سوال برای راهنمایی از هوش مصنوعی استفاده شده است.)
### سوال ششم: انتخاب روش مناسب برای بازگرداندن تغییرات
1. دستور `git restore path/to/file` این کار را انجام می‌دهد. Working tree به طبع تغییر می‌کند.
2. دستور `git restore --staged <file>` این کار را انجام می‌دهد. Staging area تغییر می‌کند.
3. با استفاده از دستور `git revert <commit_hash>` می‌توان یک کامیت جدید ایجاد کرد که تغییرات کامیت نوشته شده را revert می‌کند. بدلیل کار گروهی، حذف کامیت از تاریخچه و `force push` در صورتی که دیگران کامیت قبلی را pull کرده باشند می‌تواند باعث ناسازگاری main با کد محلی خودشان شود و فقط وضعیت را بدتر می‌کند.
(برای بخش ۳ سوال برای راهنمایی از هوش مصنوعی استفاده شده است)