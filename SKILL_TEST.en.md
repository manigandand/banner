# Task

* A product manager at Mercari has a great idea!  Let’s add banners to the Mercari app.
* We need some way of controlling which banner is displayed.
* The engineering manager has asked you to encapsulate this behavior in a library.
* Design the feature for the Banner logic, and implement it according to the specifications below.
* Technical Language
    * A banner is **expired** when the display period is over.
    * A banner’s **display period** is the duration the banner is active on the screen.
    * A banner is **active** during the display period.
* Please post any questions you have as a Github issue.


# Requirements

* You may code in any of the following languages(We use Go and PHP mainly in the company)
  * Go, PHP
* We’re looking for a well designed structure and a clean implementation, of your new banner library.
* Do not implement the API layer (e.g. HTTP, GraphQL), This should be an exportable library and is isolated from any API layer.
* Avoid using external libraries as much as possible. Try to stick with the standard library of your chosen programming language.
* Ensure that the solution is well tested, with appropriate documentation on how to run the tests.
* You also do not have to implement a data layer, feel free to stub this out with mock or test data.  However, the code should be written in a way that is easily adapted to accept a database layer in the future.


# Specifications

* Banner Display Period Conditions
    * Each banner is associated with a display period.
    * Therefore, each banner will only run for a specific period of time.
* Banner Display Preiod Rules
    * A banner is active if within the Display Period: `start time <= current time <= end time`
    * The banner's start and end time should be **timezone aware** when compared with the current time.
        * The display period starts and ends at the same moment worldwide.
        * Example: A banner starting at 03:00 UTC(+00:00), starts at 12:00 JST(+09:00).
    * Only one banner can be active at a time.
* Internal Release & QA Considerations
    * We’d like to display the banner **if the user has an internal IP address** (10.0.0.10, 10.0.0.11), even if the current time is **before** the display period of the banner.
    * **After** a banner expires, it should not be displayed again, even from an internal IP address.
    * During QA, there may be occasions where two banners are considered active.  In this case, the banner with the earlier expiration should be displayed.

