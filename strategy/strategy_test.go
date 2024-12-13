package strategy

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for os.ReadFile
type MockFileReader struct {
	mock.Mock
}

func (m *MockFileReader) ReadFile(path string) ([]byte, error) {
	args := m.Called(path)
	if bytes, ok := args.Get(0).([]byte); ok {
		return bytes, args.Error(1)
	}
	// bytesがnilの場合は空のスライスを返すか、エラーを返す
	return nil, args.Error(1)
}

// Wrap Strategy for dependency injection
type TestableStrategy struct {
	*Strategy
	FileReader func(string) ([]byte, error)
}

func (s *TestableStrategy) Load(path string) ([]byte, error) {
	bytes, err := s.FileReader(path)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func TestStrategy_Load(t *testing.T) {
	t.Run("正常系: ファイルが正常に読み込まれる", func(t *testing.T) {
		mockReader := new(MockFileReader)
		mockReader.On("ReadFile", "test.yml").Return([]byte("key: value"), nil)

		strategy := &TestableStrategy{
			Strategy:   &Strategy{},
			FileReader: mockReader.ReadFile,
		}

		data, err := strategy.Load("test.yml")
		assert.NoError(t, err, "ファイルの読み込みに成功するべきです")
		assert.Equal(t, []byte("key: value"), data, "読み込まれたデータが正しいこと")
		mockReader.AssertCalled(t, "ReadFile", "test.yml")
	})

	t.Run("異常系: ファイルが存在しない場合エラーが返る", func(t *testing.T) {
		mockReader := new(MockFileReader)
		mockReader.On("ReadFile", "not_found.yml").Return(nil, errors.New("ファイルが見つかりません"))

		strategy := &TestableStrategy{
			Strategy:   &Strategy{},
			FileReader: mockReader.ReadFile,
		}

		_, err := strategy.Load("not_found.yml")
		assert.Error(t, err, "ファイルが見つからない場合エラーが返るべきです")
		assert.Contains(t, err.Error(), "ファイルが見つかりません", "エラーメッセージが正しいこと")
		mockReader.AssertCalled(t, "ReadFile", "not_found.yml")
	})
}
