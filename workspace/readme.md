# 切片Slice

* 切片slice：切片是对数组的抽象。Go数组的长度在定义后是固定的，不可改变的。切片的长度和容量是不固定的，可以动态增加元素，切片的容量也会根据情况扩容

* 定义和初始化

  * 语法

    ```go
    var slice_var []data_type 
    var slice_var []data_type = make([]data_type, len, cap)// cap是切片容量，是make的可选参数
    var slice_var []data_type = make([]data_type, len)
    ```

    

  * 示例

* 默认值nil

* 使用

* 切片截取

* len()和cap()函数

* append()和copy()函数

* 函数传参

