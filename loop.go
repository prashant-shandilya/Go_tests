package main

import (
	"fmt"
)

func main() {
	
	now := []string{"polkadot","cabbage","carrot","potato","tomato","onion","garlic","ginger","chilli","pepper","cucumber","pumpkin","brinjal","beetroot","radish","turnip","cauliflower","cabbage","broccoli","lettuce","spinach","coriander","mint","basil","thyme","oregano","rosemary","sage","parsley","dill","lemon","lime","orange","grapefruit","pomelo","mandarin","tangerine","papaya","mango","banana","jackfruit","pineapple","watermelon","muskmelon","cantaloupe","honeydew","apple","pear","peach","plum","apricot","cherry","strawberry","blueberry","raspberry","blackberry","cranberry","currant","grape","fig","date","kiwi","avocado","pomegranate","lychee","longan","rambutan","durian","custard apple","passion fruit","guava","star fruit","dragon fruit","breadfruit","coconut","cashew","almond","walnut","pistachio","pecan","hazelnut","peanut","chestnut","macadamia","sesame","sunflower","pumpkin","flax","chia","quinoa","amaranth","buckwheat","oat","rice","wheat","barley","rye","corn","millet","sorghum","soybean","lentil","chickpea","kidney bean","black bean","pinto bean","navy bean","lima bean","broad bean","mung bean","adzuki bean","black-eyed pea","green gram","pea","carrot","beetroot","radish","turnip","cauliflower","cabbage","broccoli","lettuce","spinach","coriander","mint","basil","thyme","oregano","rosemary","sage","parsley","dill","lemon","lime","orange","grapefruit","pomelo","mandarin","tangerine","papaya","mango","banana","jackfruit","pineapple","watermelon","muskmelon","cantaloupe","honeydew","apple","pear","peach","plum","apricot","cherry",}

	for k,v := range now {
		fmt.Println(k,v)
	}

}