package highlight

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sivukhin/gopeg/extension"
)

func TestPython(t *testing.T) {
	rules, err := extension.Load(PythonTokenizer)
	require.Nil(t, err)
	highlighted, err := Highlight(`def square(value): 
    result = value**2 # dummy comment
    return result`, rules)
	require.Nil(t, err)
	require.Equal(t, `<span class="keyword">def</span> <span class="function">square</span>(<span class="identifier">value</span>): 
    <span class="identifier">result</span> = <span class="identifier">value</span>**2 <span class="comment"># dummy comment</span>
    <span class="keyword">return</span> <span class="identifier">result</span>`, highlighted)
}
