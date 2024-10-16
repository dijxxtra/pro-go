package main

// //6.
// func writeName(val struct {
// 	name, category string
// 	price          float64
// }) {
// 	fmt.Println("Name:", val.name)
// }

// //8.2
// type Product struct {
// 	name, category string
// 	price          float64
// }
// func calcTax(product *Product) {
// 	if (*product).price > 100 {
// 		(*product).price += (*product).price * 0.2
// 	}
// }
// Go следует по указателям без символа *
// func calcTax(product *Product) {
// 	if product.price > 100 {
// 		product.price += product.price * 0.2
// 	}
// }

// //8.3
// type Product struct {
// 	name, category string
// 	price          float64
// }
// func newProduct(name, category string, price float64) *Product {
// 	return &Product{name, category, price}
// }

// //8.4
// type Product struct {
// 	name, category string
// 	price          float64
// 	*Supplier
// }
// type Supplier struct {
// 	name, city string
// }

// func newProduct(name, category string, price float64, supplier *Supplier) *Product {
// 	return &Product{name, category, price - 10, supplier}
// }

// //8.4.1
// func copyProduct(product *Product) Product {
// 	p := *product
// 	s := *product.Supplier
// 	p.Supplier = &s
// 	return p
// }

// // 8.5
// type Product struct {
// 	name, category string
// 	price          float64
// }

// //8.5.1
// type Product struct {
// 	name, category string
// 	price          float64
// 	*Supplier
// }

// type Supplier struct {
// 	name, city string
// }

func main() {
	// //1. Определение и использование структуры
	// type Product struct {
	// 	name, category string
	// 	price          float64
	// }
	// // Создание структурных значений
	// kayak := Product{
	// 	name:     "Kayak",
	// 	category: "Watersports",
	// 	price:    275}
	// // Использование значения структуры
	// fmt.Println(kayak.name, kayak.category, kayak.price)
	// kayak.price = 300
	// fmt.Println("Changed price:", kayak.price)

	// // Частичное присвоение значений структуры
	// type Product struct {
	// 	name, category string
	// 	price          float64
	// }
	// kayak := Product{
	// 	name:     "Kayak",
	// 	category: "Watersports",
	// }
	// fmt.Println(kayak.name, kayak.category, kayak.price)
	// kayak.price = 300
	// fmt.Println("Changed price:", kayak.price)
	// var lifejacket Product
	// fmt.Println("Name is zero value:", lifejacket.name == "")
	// fmt.Println("Category is zero value:", lifejacket.category == "")
	// fmt.Println("Price is zero value:", lifejacket.price == 0)

	// //lifejacket := new(Product) == lifejacket:= &Product{}

	// //2. Использование позиций полей для создания значений структуры
	// type Product struct {
	// 	name, category string
	// 	price          float64
	// }
	// var kayak = Product{"Kayak", "Watersports", 275.00}
	// fmt.Println("Name:", kayak.name)
	// fmt.Println("Category:", kayak.category)
	// fmt.Println("Price:", kayak.price)

	// //3. Определение встроенных полей
	// type Product struct {
	// 	name, category string
	// 	price          float64
	// }
	// type StockLevel struct {
	// 	Product
	// 	count int
	// }
	// stockItem := StockLevel{
	// 	Product: Product{"Kayak", "Watersports", 275.00},
	// 	count:   100,
	// }
	// fmt.Println("Name:", stockItem.Product.name)
	// fmt.Println("Count:", stockItem.count)

	// //4. Сравнение значений структуры
	// type Product struct {
	// 	name, category string
	// 	price          float64
	// }
	// p1 := Product{"Kayak", "Watersports", 275.00}
	// p2 := Product{"Kayak", "Watersports", 275.00}
	// p3 := Product{"Kayak", "Boats", 275.00}
	// fmt.Println("p1 == p2:", p1 == p2)
	// fmt.Println("p1 == p3:", p1 == p3)

	// //4.1
	// type Product struct {
	// 	name, category string
	// 	price          float64
	// 	otherNames     []string
	// }
	// p1 := Product{"Kayak", "Watersports", 275.00}
	// p2 := Product{"Kayak", "Watersports", 275.00}
	// p3 := Product{"Kayak", "Boats", 275.00}
	// // Нельзя сравнивать т.к в структуре есть несравниваемое поле []string
	// fmt.Println("p1 == p2:", p1 == p2)
	// fmt.Println("p1 == p3:", p1 == p3)

	// //5. Преобразование между типами структур
	// type Product struct {
	// 	name, category string
	// 	price          float64
	// }
	// type Item struct {
	// 	name     string
	// 	category string
	// 	price    float64
	// }
	// prod := Product{"Kayak", "Watersports", 275.00}
	// item := Item{"Kayak", "Watersports", 275.00}
	// // Product(item) - преобразование в Product
	// fmt.Println("prod == item:", prod == Product(item))

	// //6. Определение анонимных типов структур
	// type Product struct {
	// 	name, category string
	// 	price          float64
	// }
	// type Item struct {
	// 	name, category string
	// 	price          float64
	// }
	// prod := Product{"Kayak", "Watersports", 275.00}
	// item := Product{"Stadium", "Soccer", 75000.00}
	// writeName(prod)
	// writeName(item)

	// //7. Создание массивов, срезов и карт, содержащих структурные значения
	// type Product struct {
	// 	name, category string
	// 	price          float64
	// }

	// type StockLevel struct {
	// 	Product
	// 	Alternate Product
	// 	count     int
	// }

	// array := [1]StockLevel{
	// 	{
	// 		Product:   Product{"Kayak", "Watersports", 275.00},
	// 		Alternate: Product{"Lifejacket", "Watersports", 48.95},
	// 		count:     100,
	// 		// Product{"Kayak", "Watersports", 275.00},
	// 		// Product{"Lifejacket", "Watersports", 48.95},
	// 		// 100,
	// 	},
	// }
	// fmt.Println("Array:", array[0].Product.name)
	// slice := []StockLevel{
	// 	{
	// 		Product:   Product{"Kayak", "Watersports", 275.00},
	// 		Alternate: Product{"Lifejacket", "Watersports", 48.95},
	// 		count:     100,
	// 	},
	// }
	// fmt.Println("Slice:", slice[0].Product.name)

	// kvp := map[string]StockLevel{
	// 	"kayak": {
	// 		Product:   Product{"Kayak", "Watersports", 275.00},
	// 		Alternate: Product{"Lifejacket", "Watersports", 48.95},
	// 		count:     100,
	// 	},
	// }
	// fmt.Println("Map:", kvp["kayak"].Product.name)

	// //8. Понимание структур и указателей
	// type Product struct {
	// 	name, category string
	// 	price          float64
	// }
	// p1 := Product{
	// 	name:     "Kayak",
	// 	category: "Watersports",
	// 	price:    275.00,
	// }
	// p2 := p1 // копирует значение
	// p1.name = "Original Kayak"
	// fmt.Println("P1:", p1.name) // P1: Original Kayak
	// fmt.Println("P2:", p2.name) // P2: Kayak

	// //8.1
	// type Product struct {
	// 	name, category string
	// 	price          float64
	// }
	// p1 := Product{
	// 	name:     "Kayak",
	// 	category: "Watersports",
	// 	price:    275.00,
	// }
	// p2 := &p1
	// p1.name = "Original Kayak"
	// fmt.Println("P1:", p1.name)
	// fmt.Println("P2:", (*p2).name)

	// //8.2 Понимание удобного синтаксиса указателя структуры
	// kayak := Product{"Kayak", "Watersports", 275.00}
	// calcTax(&kayak)
	// fmt.Println("Name:", kayak.name, "Category:", kayak.category, "Price:", kayak.price)

	// //8.2 Понимание указателей на значения
	// kayak := &Product{"Kayak", "Watersports", 275.00}
	// calcTax(kayak)
	// fmt.Println("Name:", kayak.name, "Category:", kayak.category, "Price:", kayak.price)

	// //8.3 Понимание функций конструктора структуры
	// products := [2]*Product{
	// 	newProduct("Kayak", "Watersports", 275),
	// 	newProduct("Hat", "Skiing", 42.50),
	// }
	// for _, p := range products {
	// 	fmt.Println("Name:", p.name, "Category:", p.category, "Price:", p.price)
	// }

	// //8.4 Использование типов указателей для полей структуры
	// acme := &Supplier{"Acme Co", "New York"}
	// products := [2]*Product{
	// 	newProduct("Kayak", "Watersports", 275, acme),
	// 	newProduct("Hat", "Skiing", 42.5, acme),
	// }
	// for _, p := range products {
	// 	fmt.Printf("Name: %s, Category: %s, Price: %.1f, Supplier: %s, %s\n", p.name, p.category, p.price, p.Supplier.name, p.Supplier.city)
	// }

	// //8.4 Общие сведения о копировании поля указателя
	// acme := &Supplier{"Acme Co", "New York"}
	// p1 := newProduct("Kayak", "Wotersports", 275, acme)
	// p2 := *p1

	// p1.name = "Original Kayak"
	// p1.Supplier.name = "BoatCo"

	// for _, p := range []Product{*p1, p2} {
	// 	fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name, p.Supplier.city)
	// }

	// //8.4.1
	// acme := &Supplier{"Acme Co", "New York"}
	// p1 := newProduct("Kayak", "Watersports", 275, acme)
	// p2 := copyProduct(p1)

	// p1.name = "Original Kayak"
	// p1.Supplier.name = "BoatCo"
	// for _, p := range []Product{*p1, p2} {
	// 	fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name)
	// }

	// // 8.5 Понимание нулевого значения для структур и указателей на структуры
	// var prod Product
	// var prodPtr *Product
	// fmt.Println("Value:", prod.name, prod.category, prod.price) //""""0
	// fmt.Println("Pointer:", prodPtr) // nil

	// //8.5.1
	// var prod Product
	// var prodPtr *Product
	// fmt.Println("Value:", prod.name, prod.category, prod.price, prod.Supplier.name) // Ошибка т.к Supplier указатель, а он равен nil
	// fmt.Println("Pointer:", prodPtr)                                                // nil

	// // var prod Product = Product{Supplier: &Supplier{}}
}
