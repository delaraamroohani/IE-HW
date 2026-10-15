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
