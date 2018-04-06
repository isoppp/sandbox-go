echo ------------------------------
echo "go run main.go -n 1 dir dir2"
echo ------------------------------
go run main.go -n 1 dir dir2

echo ------------------------------
echo "go run main.go -n 1 ."
echo ------------------------------
go run main.go -n 1 .

echo ------------------------------
echo "go run main.go -n 1 ./"
echo ------------------------------
go run main.go -n 1 ./

echo ------------------------------
echo "go run main.go -n 1 sample.txt"
echo ------------------------------
go run main.go -n 1 sample.txt

echo ------------------------------
echo "go run main.go -n 1 dir/*.html"
echo ------------------------------
go run main.go -n 1 dir/*.html

echo ------------------------------
echo "go run main.go -n 1 **/*.html"
echo ------------------------------
go run main.go -n 1 **/*.html

echo ------------------------------
echo "go run main.go **"
echo ------------------------------
go run main.go **

echo ------------------------------
echo "go run main.go -n 1 -s dir2"
echo ------------------------------
go run main.go -n 1 -s dir2
