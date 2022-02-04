/*
 * @lc app=leetcode.cn id=146 lang=cpp
 *
 * [146] LRU 缓存
 */

// @lc code=start
#include <map>

using namespace std;
class Node
{
private:
public:
    int key, value;
    Node *prev, *next;
    Node(int k, int v);
    ~Node();
};

Node::Node(int k, int v)
{
    key = k;
    value = v;
    prev = nullptr;
    next = nullptr;
}

Node::~Node()
{
}

class DoubleList
{
private:
    Node *head, *tail;
    int len;

public:
    DoubleList();
    void addLast(Node *x);
    void remove(Node *x);
    Node *removeFirst();
    int size();
    ~DoubleList();
};

DoubleList::DoubleList()
{
    head = new Node(0, 0);
    tail = new Node(0, 0);
    head->next = tail;
    tail->prev = head;
    len = 0;
}

void DoubleList::addLast(Node *x)
{
    tail->prev->next = x;
    x->next = tail;
    x->prev = tail->prev;
    tail->prev = x;
    len++;
}

void DoubleList::remove(Node *x)
{
    x->prev->next = x->next;
    x->next->prev = x->prev;
    len--;
    // delete x;
}

// 节点未释放，调用后在上层释放
Node *DoubleList::removeFirst()
{
    if (head->next == tail)
        return nullptr;

    Node *first = head->next;

    remove(first);
    return first;
}

int DoubleList::size()
{
    return len;
}

DoubleList::~DoubleList()
{
}

class LRUCache
{
private:
    map<int, Node *> cacheMap;
    DoubleList *cache;
    int cap;

    void makeRecently(int key);
    void addRecently(int key, int value);
    void deleteKey(int key);
    void removeLeastRecently();

public:
    LRUCache(int capacity)
    {
        cap = capacity;
        cacheMap = map<int, Node *>();
        cache = new DoubleList();
    }

    int get(int key)
    {
        if (!cacheMap.count(key))
        {
            return -1;
        }
        makeRecently(key);

        return cacheMap[key]->value;
    }

    void put(int key, int value)
    {
        if (cacheMap.count(key))
        {
            Node *x = cacheMap[key];
            x->value = value;
            cache->remove(x);
            cache->addLast(x);

            return;
        }

        if (cacheMap.size() >= cap)
        {
            Node *deleteNode = cache->removeFirst();
            cacheMap.erase(deleteNode->key);
        }

        Node *addNode = new Node(key, value);
        cache->addLast(addNode);
        cacheMap[key] = addNode;
    }
};

void LRUCache::makeRecently(int key)
{
    Node *x = cacheMap[key];
    cache->remove(x);
    cache->addLast(x);
}

void LRUCache::addRecently(int key, int value)
{
    Node *x = new Node(key, value);
    cache->addLast(x);
    cacheMap.insert(pair<int, Node *>(key, x));
}

void LRUCache::deleteKey(int key)
{
    Node *x = cacheMap[key];
    cache->remove(x);
    cacheMap.erase(key);
    delete x;
}

void LRUCache::removeLeastRecently()
{
    Node *deleteNode = cache->removeFirst();
    int key = deleteNode->key;
    cacheMap.erase(key);
    delete deleteNode;
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */
// @lc code=end

int main()
{
    LRUCache lRUCache = LRUCache(2);
    lRUCache.put(1, 1); // 缓存是 {1=1}
    lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
    lRUCache.get(1);    // 返回 1
    lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
    lRUCache.get(2);    // 返回 -1 (未找到)
    lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
    lRUCache.get(1);    // 返回 -1 (未找到)
    lRUCache.get(3);    // 返回 3
    lRUCache.get(4);    // 返回 4
}
