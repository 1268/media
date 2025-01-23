# kanopy

~~~
url = https://www.kanopy.com/en/product/13808102?vp=hclib
monetization = FREE
country = United States
~~~

- https://play.google.com/store/apps/details?id=com.kanopy
- https://play.google.com/store/apps/details?id=com.kanopy.tvapp

here is the layout:

~~~
episode   .children   nil
.parent
season    .children   episodes
.parent
show      .children   seasons
.parent
nil
~~~
