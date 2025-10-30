# Компилятор C++
CXX = g++
# Флаги компиляции: стандарт C++11, все предупреждения, отладочная информация
CXXFLAGS = -std=c++11 -Wall -Wextra -g

# Имя итоговой программы
TARGET = lab2

# Директория для скомпилированных файлов
BUILD_DIR = ../../build/lab2

# Пути к заголовочным файлам (папка common/include и текущая папка lab2)
INCLUDES = -I../../common/include -I.

# Исходные файлы из текущей лабораторной (labs/lab2)
LAB_SOURCES = main.cpp \
              expressionSolver.cpp \
              patternMatcher.cpp \
              avlChecker.cpp \
              foldingHasher.cpp \
              lruCache.cpp

# Исходные файлы из common, от которых зависят задания lab2
COMMON_SOURCES = ../../common/src/stack.cpp \
                 ../../common/src/array.cpp \
                 ../../common/src/binaryTree.cpp

# Автоматически генерируем пути к объектным файлам (.o)
LAB_OBJECTS = $(patsubst %.cpp, $(BUILD_DIR)/%.o, $(LAB_SOURCES))
COMMON_OBJECTS = $(patsubst ../../common/src/%.cpp, $(BUILD_DIR)/%.o, $(COMMON_SOURCES))
OBJECTS = $(LAB_OBJECTS) $(COMMON_OBJECTS)

# Правило по умолчанию: собрать всё
all: $(BUILD_DIR) $(TARGET)

# Правило для линковки (сборки) исполняемого файла из объектных файлов
$(TARGET): $(OBJECTS)
	$(CXX) $(CXXFLAGS) $(OBJECTS) -o $(BUILD_DIR)/$(TARGET)

# Правило для создания директории сборки
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Правило для компиляции .cpp из папки labs/lab2
$(BUILD_DIR)/%.o: %.cpp
	$(CXX) $(CXXFLAGS) $(INCLUDES) -c $< -o $@

# Правило для компиляции .cpp из падки common/src
$(BUILD_DIR)/%.o: ../../common/src/%.cpp
	$(CXX) $(CXXFLAGS) $(INCLUDES) -c $< -o $@

# Правило для очистки
clean:
	rm -rf $(BUILD_DIR)