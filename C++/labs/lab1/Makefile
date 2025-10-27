# Используем CXX для C++ компилятора
CXX = g++
# Флаги для компилятора. Добавил -g для отладки
CXXFLAGS = -std=c++11 -Wall -Wextra -g

# Имя исполняемого файла
TARGET = dbms

# Директория для сборки. ../../build/lab1, чтобы сборки разных лаб не смешивались
BUILD_DIR = ../../build/lab1

INCLUDES = -I../../common/include

LAB_SOURCES = main.cpp

COMMON_SOURCES = ../../common/src/array.cpp \
                 ../../common/src/binaryTree.cpp \
                 ../../common/src/dlinkedList.cpp \
                 ../../common/src/linkedList.cpp \
                 ../../common/src/queue.cpp \
                 ../../common/src/stack.cpp \
                 ../../common/src/set.cpp \
				 ../../common/src/hashTable.cpp

# Объединяем списки
SOURCES = $(LAB_SOURCES) $(COMMON_SOURCES)

# Автоматически генерируем пути к ОБЪЕКТНЫМ файлам (.o), которые будут лежать в BUILD_DIR
OBJECTS = $(patsubst %.cpp, $(BUILD_DIR)/%.o, $(LAB_SOURCES)) \
          $(patsubst ../../common/src/%.cpp, $(BUILD_DIR)/%.o, $(COMMON_SOURCES))

# --- ПРАВИЛА СБОРКИ ---

# Правило по умолчанию: создать папку для сборки и собрать исполняемый файл
all: $(BUILD_DIR) $(TARGET)

# Правило для создания исполняемого файла из объектных файлов
$(TARGET): $(OBJECTS)
	$(CXX) $(CXXFLAGS) $(OBJECTS) -o $(BUILD_DIR)/$(TARGET)

# Правило для создания папки для сборки
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Правило для компиляции .cpp файлов из ТЕКУЩЕЙ папки в .o файлы в BUILD_DIR
$(BUILD_DIR)/%.o: %.cpp
	$(CXX) $(CXXFLAGS) $(INCLUDES) -c $< -o $@

# Правило для компиляции .cpp файлов из COMMON папки в .o файлы в BUILD_DIR
$(BUILD_DIR)/%.o: ../../common/src/%.cpp
	$(CXX) $(CXXFLAGS) $(INCLUDES) -c $< -o $@

# Правило для очистки
clean:
	rm -rf $(BUILD_DIR)