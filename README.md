# calculator_test
Тестовое задание. Калькулятор римских и арабских чисел.

<p><b><h2>Описание задачи</h2></b></p>
<p>Создай консольное приложение «Калькулятор». Приложение должно читать из консоли введенные пользователем строки, числа, арифметические операции, проводимые между ними, и выводить в консоль результат их выполнения.
Калькулятор можно реализовать обычными функциями либо использовать структуру с методами, здесь это не принципиально.</p>

<p><b><h3>Требования:</h3></b></p>
<p><ul><li>Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: a + b, a - b, a * b, a / b. Данные передаются в одну строку (смотри пример ниже). Решения, в которых каждое число и арифметическая операция передаются с новой строки, считаются неверными.</li>
<li>Калькулятор умеет работать как с арабскими (1,2,3,4,5…), так и с римскими (I,II,III,IV,V…) числами.</li>
<li>Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более. На выходе числа не ограничиваются по величине и могут быть любыми.</li>
<li>Калькулятор умеет работать только с целыми числами.</li>
<li>Калькулятор умеет работать только с арабскими или римскими цифрами одновременно, при вводе пользователем строки вроде 3 + II калькулятор должен указать на ошибку и прекратить работу.</li>
<li>При вводе римских чисел ответ должен быть выведен римскими цифрами, соответственно, при вводе арабских — ответ ожидается арабскими.</li>
<li>При вводе пользователем не подходящих чисел приложение выводит ошибку в терминал и завершает работу.</li>
<li>При вводе пользователем строки, не соответствующей одной из вышеописанных арифметических операций, приложение выводит ошибку и завершает работу.</li>
<li>Результатом операции деления является целое число, остаток отбрасывается.</li>
<li>Результатом работы калькулятора с арабскими числами могут быть отрицательные числа и ноль. Результатом работы калькулятора с римскими числами могут быть только положительные числа, если результат работы меньше единицы, программа должна указать на исключение.</li></ul></p>
<p><h3>Пример работы программы:</h3></p>
<pre>
Input:
1 + 2<br>

Output:
3

Input:
VI / III

Output:
II

Input:
I - II

Output:
Вывод ошибки, так как в римской системе нет отрицательных чисел.

Input:
I + 1

Output:
Вывод ошибки, так как используются одновременно разные системы счисления.

Input:
1

Output:
Вывод ошибки, так как строка не является математической операцией.

Input:
1 + 2 + 3

Output:
Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).
</pre>
