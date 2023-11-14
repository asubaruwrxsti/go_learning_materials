# A tour of GO
:::info
[Github Repo](https://github.com/benoitvallon/a-tour-of-go/tree/master)
:::
* Pointers
    ```
    Go has pointers. A pointer holds the memory address of a value.

    The type *T is a pointer to a T value. Its zero value is nil.

    var p *int
    The & operator generates a pointer to its operand.

    i := 42
    p = &i
    The * operator denotes the pointer's underlying value.

    fmt.Println(*p) // read i through the pointer p
    *p = 21         // set i through the pointer p
    This is known as "dereferencing" or "indirecting".

    Unlike C, Go has no pointer arithmetic.
    ```
* Slices
    * Slices are like references to arrays. A slice does not store any data, it just describes a section of an underlying array.
    * Slice literals are like an array literal without the length.
        * This is an array literal
            ```
            [3]bool{true, true, false}
            ```
        * And this creates the same array as above, then builds a slice that references it:
            ```
            []bool{true, true, false}
            ```
    * Slice defaults
        * These expressions are equivalent
            ```
            a[0:10]
            a[:10]
            a[0:]
            a[:]
            ```
    * Slice length and capacity
        * A slice has both a length and a capacity.
        * The length of a slice is the number of elements it contains.
        * The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
    * Nil slices
        * A nil slice has a length and capacity of 0 and has no underlying array.
    * Creating a slice with **make**
        * Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
        * The make function allocates a zeroed array and returns a slice that refers to that array:
        ```
            a := make([]int, 5)  // len(a)=5
        ```
        * To specify a capacity, pass a third argument to make:
        ```
            b := make([]int, 0, 5) // len(b)=0, cap(b)=5
            b = b[:cap(b)] // len(b)=5, cap(b)=5
            b = b[1:]      // len(b)=4, cap(b)=4
        ```
    * Appending to a slice
        ```
            func append(s []T, vs ...T) []T
        ```
        * The first parameter s of append is a slice of type T, and the rest are T values to append to the slice. The returned slice will point to the newly allocated array.

* Range
    * The range form of the for loop iterates over a slice or map.
    * When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
        ```
            var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

            func main() {
                for i, v := range pow {
                    fmt.Printf("2**%d = %d\n", i, v)
                }
            }
        ```
    * You can skip the index or value by assigning to _.
        ```
            for i, _ := range pow
            for _, value := range pow
        ```
    * If you only want the index, you can omit the second variable.
        ```
            for i := range pow
        ```
    * Sample:
        ```
            func main() {
                pow := make([]int, 10)
                for i := range pow {
                    pow[i] = 1 << uint(i) // == 2**i
                }
                for _, value := range pow {
                    fmt.Printf("%d\n", value)
                }
            }
        ```
* Maps
    * A map maps keys to values
    * The make function returns a map of the given type, initialized and ready for use.
        ```
            type Vertex struct {
            Lat, Long float64
            }

            var m map[string]Vertex

            func main() {
                m = make(map[string]Vertex)
                m["Bell Labs"] = Vertex{
                    40.68433, -74.39967,
                }
                fmt.Println(m["Bell Labs"])
            }
        ```
    * Map literals are like struct literals, but the keys are required.
        ```
            type Vertex struct {
                Lat, Long float64
            }

            var m = map[string]Vertex{
                "Bell Labs": Vertex{
                    40.68433, -74.39967,
                },
                "Google": Vertex{
                    37.42202, -122.08408,
                },
            }

            func main() {
                fmt.Println(m)
            }
        ```
    * If the top-level type is just a type name, you can omit it from the elements of the literal.
        ```
            type Vertex struct {
                Lat, Long float64
            }

            var m = map[string]Vertex{
                "Bell Labs": {40.68433, -74.39967},
                "Google":    {37.42202, -122.08408},
            }

            func main() {
                fmt.Println(m)
            }
        ```
    * Mutating Maps
        * Insert or update an element in map m:
            ```
            m[key] = elem
            ```
        * Retrieve an element:
            ```
            elem = m[key]
            ```
        * Delete an element:
            ```
            delete(m, key)
            ```
        * Test that a key is present with a two-value assignment:
            ```
            elem, ok = m[key]
            ```
* Function values
    * Function values may be used as function arguments and return values.
        ```
            func compute(fn func(float64, float64) float64) float64 {
                return fn(3, 4)
            }

            func main() {
                hypot := func(x, y float64) float64 {
                    return math.Sqrt(x*x + y*y)
                }
                fmt.Println(hypot(5, 12))

                fmt.Println(compute(hypot))
                fmt.Println(compute(math.Pow))
            }
        ```
* Function closures
    * A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
    * The adder function returns a closure. Each closure is bound to its own sum variable.
        ```
            func adder() func(int) int {
                sum := 0
                return func(x int) int {
                    sum += x
                    return sum
                }
            }

            func main() {
                pos, neg := adder(), adder()
                for i := 0; i < 10; i++ {
                    fmt.Println(
                        pos(i),
                        neg(-2*i),
                    )
                }
            }
        ```
* Methods and Classes
    * A method is a function with a special receiver argument. The receiver appears in its own argument list between the func keyword and the method name. In this example, the Abs method has a receiver of type Vertex named v.
        ```
            type Vertex struct {
                X, Y float64
            }

            func (v Vertex) Abs() float64 {
                return math.Sqrt(v.X*v.X + v.Y*v.Y)
            }

            func main() {
                v := Vertex{3, 4}
                fmt.Println(v.Abs())
            }
        ```
    * Methonds are functions with a receiver argument
        * Abs written as a regular function with no change in functionality.
        ```
            type Vertex struct {
                X, Y float64
            }

            func Abs(v Vertex) float64 {
                return math.Sqrt(v.X*v.X + v.Y*v.Y)
            }

            func main() {
                v := Vertex{3, 4}
                fmt.Println(Abs(v))
            }
        ```
        * You can declare a method on non-struct types, too.
        * You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
            ```
                type MyFloat float64

                func (f MyFloat) Abs() float64 {
                    if f < 0 {
                        return float64(-f)
                    }
                    return float64(f)
                }

                func main() {
                    f := MyFloat(-math.Sqrt2)
                    fmt.Println(f.Abs())
                }
            ```
