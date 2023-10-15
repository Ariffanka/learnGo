package helper

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"fmt"
)

//mencoba benchmark table
func BenchmarkTable(b *testing.B){
	benchmarks := []struct{
		name string
		request string
	}{
		{
			name:"arip",
			request: "arip",
		},
		{
			name:"panka",
			request: "panka",
		},
	}

	for _,  benchmark := range benchmarks {
		b.Run(benchmark.name, func (b *testing.B){
			for i:=0; i < b.N; i++{
				Hello(benchmark.request)
			}
		})
	}
}

//mencoba sub benchmark testing
func BenchmarkSub(b *testing.B){

	b.Run("arip", func (b *testing.B){
		for i:=0; i < b.N; i++{
			Hello("arip")
		}
	})

	b.Run("panka", func (b *testing.B){
		for i:=0; i < b.N; i++{
			Hello("panka")
		}
	})
}

//mencoba benchmark testing
func BenchmarkHello(b *testing.B){
	for i:=0; i < b.N; i++{
		Hello("arip")
	}
}

func BenchmarkHelloLagi(b *testing.B){
	for i:=0; i < b.N; i++{
		Hello("panka")
	}
}
//mencoba table test(agar efesien dalam menggunakan sub test)
func TestTableHello(t *testing.T){

	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name:"arip",
			request:"arip",
			expected:"Hello arip",
		},

		{
			name:"panka",
			request:"panka",
			expected:"Hello panka",
		},
	}
	
	for _, test := range tests{
		t.Run(test.name, func(t *testing.T){
			result := Hello(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

//mencoba sub test
func TestSubTest(t *testing.T){
	t.Run("arip", func (t *testing.T){
		result := Hello("arip")
		require.Equal(t, "Hello arip", result, "Result must be 'Hello arip' ")
	})

	t.Run("panka", func (t *testing.T){
		result := Hello("panka")
		require.Equal(t, "Hello panka", result, "Result must be 'Hello panka' ")
	})
}

//mencoba after dan before(bukan after berfore  artian sebenarnya)
func TestMain(m *testing.M){
	//before ceritanya
	fmt.Println("BEFORE TEST")

	m.Run()

	//after ceritanya
	fmt.Println("AFTER TEST")
}

//mencoba skip test
func TestSkip(t *testing.T){
	if runtime.GOOS=="windows"{
		t.Skip("Can not run on Windows")
	}

	result := Hello("Arip")
	require.Equal(t, "Hello Arip", result, "must be run 'Hello Arip' ")
}

//mencoba require dari testify
func TestHelloRequire(t *testing.T){
	result := Hello("arip")
	require.Equal(t, "Hello arip", result, "Result must be 'Hello Arip' ")
	fmt.Println("TestHello with Require done")
}

//mencoba assertion dari testify
func TestHelloAssertion(t *testing.T){
	result := Hello("arip")
	assert.Equal(t, "Hello arip", result, "Result must be 'Hello Arip' ")
	fmt.Println("TestHello with Assertion done")
}

//penanganan error unit test
func TestHello(t *testing.T){
	result := Hello("arip")

	if result != "Hello arip"{
		// t.FailNow()
		t.Error("must be Hello (name)")
	}
	fmt.Println("Test Hello done")
}

//penanganan error unit test yang kedua
func TestHelloLagi(t *testing.T){
	result := Hello("arip")

	if result != "Hello arip"{
		// t.Fail()
		t.Fatal("must be Hello (name)")
	}
	fmt.Println("Test Hello done")
}