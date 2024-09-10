# gogeo

### Installation 
```cmd
go get github.com/rajabit/gogeo
```

### Test 
```cmd
go test
```

### Usage

Calculate distance

```go
point01 := gogeo.NewPoint(36.26729, 46.3783142)
point02 := gogeo.NewPoint(36.256909, 46.265876)
distance := point01.DistanceInMeter(point02)
```
Circle Contains
```go
point := gogeo.NewPoint(31.40028516, 51.5928680)
circlePolygon := gogeo.NewPolygon([][]float32{{51.59366908460324, 31.40037506998533}})
contain := circlePolygon.CircleContains(point, 123.123 /*radius*/)
```
Rectangle Contains
```go
point := gogeo.NewPoint(32.708254, 51.718520)
rectanglePolygon := gogeo.NewPolygon([][]float32{{51.718521, 32.708254}, {51.718521, 32.708796}, {51.719561, 32.708796}, {51.719561, 32.708254}, {51.718521, 32.708254}})
contain := rectanglePolygon.RectangleContains(point)
```
Polygon Contains
```go
point := gogeo.NewPoint(32.671763,51.676791)
polygon := gogeo.NewPolygon([][]float32{{51.676791, 32.671763}, {51.676152, 32.671736}, {51.675557, 32.671745}, {51.675407, 32.671546}, {51.675246, 32.671257}, {51.675289, 32.671045}, {51.675546, 32.670774}, {51.675766, 32.670435}, {51.676673, 32.670286}, {51.676925, 32.670516}, {51.676952, 32.670909}, {51.676919, 32.671221}, {51.676791, 32.671763}})
contain := polygon.RectangleContains(point)
```