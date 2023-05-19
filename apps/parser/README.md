


constrains
- effects are valid for bites and spokenText
- **40 s sound limit**

filters are:
{1} room echo {2} hall echo {3} outside echo {4} pitch down {5} pitch up {6} telephone {7} muffled {8} quieter {9} ghost {10} chorus {11} slow down {12} speed up

forsen 

"voice+colon".  
Example donation: "forsen: oh shit I'm sorry trump: sorry for what?"



https://github.com/CorentinJ/Real-Time-Voice-Cloning


Filters can be used putting the filter number in curly brackets. Can cancel the last filter added with a decimal point in curly brackets i.e. {.} 

sound bites can be added with square brackets and the number of the sound bite inside the brackets  i.e[1]

All filters except {1}{2} and {9} can stack. 

`{2}weskeru:why are we in this cave forsan, why. forsen:just walk over there.[49][65]weskeru: you missed[65]missed again{.}`

`{2}[1][1]lj:Alert, patient gamma 32 is loose, all patients please remain in your cells.[1][1]{.}`

`
`[87][87][81]{6}weskeru:I don't know who you are. I don't know what you want. If you are looking for a donation I can tell you I don't have money, but what I do have are a very particular set of sniping skills. Skills I have acquired over a very long sniping career. Skills that make me a nightmare for streamers like you.{.}`

`{7}[76]forsen:help us, chat, he is an imposter. nani: we are trapped chat, only, you, can save, us. [76], help us.`

filters are:

{1} room echo
{2} hall echo
{3} outside echo
{4} pitch down
{5} pitch up
{6} telephone
{7} muffled
{8} quieter
{9} ghost
{10} chorus
{11} slow down
{12} speed up


```go

type audio struct {
	chunks []Chunk
}

type Chunk[C Segment] struct {
	filters:[]Filter
	segments:[]C
	voice: string
}

type Filter struct {
	id int
}

type SoundBite strut {
	id int
}

type SpokenText  string

func (* SpokenText)

type Segment interface {
    SoundBite | SpokenText
}

type SoundBite struct {
	id int
}

`{7}[76]forsen:help us, chat, he is an imposter. nani: {11} we are trapped chat, only, you, can save, us. [76], help us.`

audio [
	{
		effects:[7]
		segments:[
			bite(76),
			spokenText{
				voice: forsen
				text: "help us, chat, he is an imposter."
			}
		]
	}
	{
		effects:[7, 11]
		segments:[			
			spokenText{
				voice: nani
				text: "we are trapped chat, only, you, can save, us."
			},
			bite(76),
			spokenText{
				voice: nani
				text: "help us."
			}
		]
	}
]

```

an audio is build by multiple audio segments 
every change of voice its a new audio segment
every time we add a new sound effect we add a new audio segment
	this applies even if the voice has not changed i.e 
carry all the prev sound effects to the next segment 



```go
	
{7}[76]forsen:help us, chat, he is an imposter. nani: {11} we are trapped chat, only, you,{3} can save, us. [76], help us.

audio [
	{
		effects:[7]
		segments:[
			bite(76),
			spokenText{
				voice: forsen
				text: "help us, chat, he is an imposter."
			}
		]
	}
	{
		effects:[7, 11]
		segments:[			
			spokenText{
				voice: nani
				text: "we are trapped chat, only, you,"
			},
			bite(76)
		]
	},
	{
		effects:[7, 11, 3]
		segments:[			
			spokenText{
				voice: nani
				text: "can save, us."
			},
			bite(76),
			spokenText{
				voice: nani
				text: "help us."
			}
		]
	}
]

```


an alternative for a segment could be 

```go 
{
		effects:[7, 11, 3]
		voice: nani 
		segments:[						
			spokenText: "can save, us.",
			bite(76),		
			spokenText: "help us.",			
		]
	}
we just add the voice if the segment has it 

{7}[76][71][74]
the prev example doesnt have any voices or text, just sound effects and bites, so the segment doest need any voices, in other words the voices are just applied for spokenText 
{
	effects:[7],
	segments:[
		bite(76),
		bite(71),
		bite(74),
	]	
}

```
