# Online-Store

### Problem Analysis
I think the accident may occur because the system didn't check quantity of product when customer checkout, and paid. Although the system has handled it, the system must reduce quantity of product in database. And I think, in database we must save the quantity of available and ordered. So that, Order Processing department can know how many items that ordered (this means the product is on checkout or paid for).

### Solution Idea
Based on my analysis above, I have some idea to solve this problem.

1. When checkout, system must be reduce the available quantity and increase ordered quantity
2. System must check if pending payment was expired. If yes, the order must be canceled and move quantity product from quantity ordered to available quantity
3. When user paid order, quantity product remains in ordered quantity column
