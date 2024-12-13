package config

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStrategy struct {
	mock.Mock
}

func (m *MockStrategy) Load(filePath string) ([]byte, error) {
	args := m.Called(filePath)
	return nil, args.Error(0)
}

func (m *MockStrategy) Unmarshal(in []byte, out any) error {
	args := m.Called(out)
	return args.Error(0)
}

func TestConfig_SetStrategy(t *testing.T) {
	t.Run("Strategy を正しく設定できる", func(t *testing.T) {
		mockStrategy := new(MockStrategy)
		c := NewConfig()
		c.SetStrategy(mockStrategy)

		assert.Equal(t, mockStrategy, c.Strategy, "Strategy が正しく設定されていません。")
	})
}

func TestConfig_SetFilePath(t *testing.T) {
	t.Run("FilePath を正しく設定できる", func(t *testing.T) {
		filePath := "./config/test.yml"
		c := NewConfig()
		c.SetFilePath(filePath)

		assert.Equal(t, filePath, c.FilePath, "FilePath が正しく設定されていません。")
	})
}

func TestConfig_Unmarshal(t *testing.T) {
	t.Run("正常系: ファイルのロードと Unmarshal が成功する", func(t *testing.T) {
		mockStrategy := new(MockStrategy)
		filePath := "./config/test.yml"
		output := map[string]string{}

		mockStrategy.On("Load", filePath).Return(nil)
		mockStrategy.On("Unmarshal", &output).Return(nil)

		c := NewConfig().SetStrategy(mockStrategy).SetFilePath(filePath)
		err := c.Unmarshal(&output)

		assert.NoError(t, err, "Unmarshal 中にエラーが発生しました。")
		mockStrategy.AssertCalled(t, "Load", filePath)
		mockStrategy.AssertCalled(t, "Unmarshal", &output)
	})

	t.Run("異常系: Load がエラーを返す", func(t *testing.T) {
		mockStrategy := new(MockStrategy)
		filePath := "./config/invalid.yml"

		mockStrategy.On("Load", filePath).Return(errors.New("ファイルが見つかりません"))

		c := NewConfig().SetStrategy(mockStrategy).SetFilePath(filePath)
		err := c.Unmarshal(map[string]string{})

		assert.Error(t, err, "Load に失敗した場合、エラーが発生するべきです。")
		mockStrategy.AssertCalled(t, "Load", filePath)
		mockStrategy.AssertNotCalled(t, "Unmarshal")
	})

	t.Run("異常系: Unmarshal がエラーを返す", func(t *testing.T) {
		mockStrategy := new(MockStrategy)
		filePath := "./config/test.yml"
		output := map[string]string{}

		mockStrategy.On("Load", filePath).Return(nil)
		mockStrategy.On("Unmarshal", &output).Return(errors.New("Unmarshal エラー"))

		c := NewConfig().SetStrategy(mockStrategy).SetFilePath(filePath)
		err := c.Unmarshal(&output)

		assert.Error(t, err, "Unmarshal に失敗した場合、エラーが発生するべきです。")
		mockStrategy.AssertCalled(t, "Load", filePath)
		mockStrategy.AssertCalled(t, "Unmarshal", &output)
	})
}
