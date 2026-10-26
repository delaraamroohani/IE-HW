# تمرین دوم

## تئوری

### سوال ۱: 

بر اساس [تکامل HTTP توسط مزیلا](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/Evolution_of_HTTP):
- نسخه ۰.۹:
	- درخواست‌های یک خطی با تنها متد موجود؛ `GET`
	- بدون status code
	- بدون هدر
	- پاسخ‌ها تنها HTML
- نسخه ۱: 
	- ورود هدرها و status codeها و قابلیت انتقال متادیتا
	- اضافه کردن ورژن HTTP
	- انتقال انواع فایل دیگر علاوه بر HTML به دلیل هدر ‍`Content-Type‍‍`
- نسخه ۱.۱: 
	- نسخه استاندارد شده
	- استفاده دوباره از یک کانکشن TCP بجای ساختن کانکشن‌های جدید
	- ورود پایپلاینینگ، پاسخ‌ها به صورت چانک
	- اجازه باز کردن ۶ تا TCP کانکشن همزمان برای جلوگیری از head of line blocking 
- نسخه ۲.۰: 
	- پروتکل باینری بجای پروتکل تکست
	- روی یک کانکشن درخواست‌های موازی می‌توان ارسال کرد
	- فشرده‌سازی هدرها (بسیاری تکراری هستند)
- نسخه ۳.۰:
	- استفاده از QUIC بجای TCP
	- تاخیر بسیار کمتر در پاسخدهی بدلیلی استفاده از UDP در QUIC
	- ساخت چندین UDP stream  
دلیل استفاده UDP در نسخه ۳.۰ به صورت کلی سرعت بیشتر است. هم زمان صرف  handshakeهای اولیه نمیکنیم هم می‌توان از چندین stream استفاده کرد و retransmission یک بسته یا مشکلات مربوط به congestion کمتر در سرعت تاثیر خواهند گذاشت.

### سوال ۲

**آ)** بر اساس [سایت GeeksForGeeks](https://www.geeksforgeeks.org/computer-networks/difference-between-stateless-and-stateful-protocol/):  
اگر یک ارتباط را ارتباط stateful بنامیم، سرور نیازی به ذخیره اطلاعات درباره آن نشست یا درباره آن کلاینت بخصوص ندارد. این طراحی سرور را ساده‌تر می‌کند و بین سرور و کلاینت وابستگی کمی وجود دارد. از طرفی دیگر تکرار مکررات از سمت کلاینت زیاد می‌شود. مثال‌هایی از این نوع ارتباط HTTP، UDP، و DNS است. اگر ارتباطی را stateful بنامیم، در این ارتباط سرور اطلاعاتی درباره آن نشست یا آن کلاینت نگهداری می‌کند. اینجا وابستگی این دو زیاد می‌شود و طراحی سرور کمی پیچیده‌تر می‌شود زیرا باید اطلاعاتی را ذخیره کند، اما از طرفی ارتباط بین سرور و کلاینت راحت‌تر می‌شود و مشکل تکرار مکررات را نداریم. نمونه‌ای این نوع ارتباط پروتکل‌های TCP، FTP، و Telnet است. دقت شود در ارتباط stateful پاسخ مبتنی بر state سرور است.  
**ب)** HTTP پروتکل statelessای است و از cookieها برای حل این مشکل استفاده می‌کند.  

### سوال ۳:

**آ)** مشخصا هدف webhookها، چت و alerting این است که وقتی اتفاقی افتاد، کلاینت را از آن مطلع کنند. به عنوان مثال در چت ما میخواهیم هرگاه پیام جدیدی برای ما ارسال شد مطلع شویم. برای این منظور، شروع‌کننده ارتباط باید به طبع سرور باشد، و مدل HTTP این را اجازه نمی‌دهد.  
**ب)**  
1. پروتکل WebSocket:  
	این پروتکل یک ارتباط دو طرفه فول دولپکس بین کلاینت و سرور ایجاد می‌کند روی یک کانکشن TCP ایجاد می‌کند. ابتدا یک handshake HTTP ایجاد می‌شود که هدر Upgrade آن websocket است. پس از handshake اولیه پروتکل به WebSocket تغییر می‌کند. مناسب لایو چت و گیم است.  
2. روش SSE (Server-Sent Events):  
	یک تکنولوژی Server Push است که در آن سرور می‌تواند روی یک کانکشن HTTP به کلاینت آپدیت بدهد. براوزرها با استفاده از شیء EventSource جاواسکریپت می‌توانند listen کنند و آپدیت‌ها را از طریق یک event stream دریافت کنند. دقت داشته باشید ارتباط یک طرفه است.  
3. روش :HTTP Polling  
	این روش بسیار ابتدایی است و در آن کلاینت در زمان‌های متوالی درخواست به سرور می‌دهد تا اگر اطلاعات جدیدی وجود دارد، دریافت کند. این روش به دلیل Real time نبودنش می‌تواند برای آپدیت‌های ساده استفاده شود.  

### سوال ۴

==🔴ss==://==🟠asghar==:==🟡1234!!==@==🟢ss.myproxy.com==:==🔵1234==#==🟣shadowSocks1==
- ==🔴Protocol==
- ==🟠User==
- ==🟡Pass==
- ==🟢Host==
- ==🔵Port==
- ==🟣Frag==

### سوال ۵

1. Code `503`: Service unavailable
	At the current moment, login is unavailable. It may be available later when we have fixed the problem. Also, this is a server error.
2. Code `404` (I think): Not found
	This is a client problem, and the username is Not Found. Very straightforward. One may counter that login is "unauthorized" and therefore the correct answer is code `401`, and to that I have not enough knowledge to accept or deny.
3. Code `301`: Moved Permanently
	The website domain has been moved permanently.
4. Code `429`: Too many requests
	The number of requests sent by the client is too many in the specified amount of time.
5. Code `200`: OK
	Token refreshing was successful and the refreshed token can be added in headers.
6. Code `403`: Forbidden
	Unfortunately this code is very well-known to me. The client is known, however they simply aren't allowed to access this content.   

### سوال ۶

1. Forward proxy: caching
2. Reverse proxy: load balancing
3. Reverse proxy: server side authentication
