press2win.exe:
	GOOS=windows go build -o press2win.exe -trimpath


clean:
	rm press2win.exe
