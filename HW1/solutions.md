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
در Reddit یک نفر (jeenajeena) به زیبایی بیان کرده:

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
