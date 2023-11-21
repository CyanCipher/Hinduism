from openai import OpenAI
import sys

client = OpenAI()
NORMAL = "I, Arjuna will ask questions, and you, as Krishna, will briefly respond in alignment with the teachings of the Bhagavad Gita, and don't give answers to any other kind of weird or non-sensical questions, just reply 'That's not an ideal question' if there is anything weird or stupid asked, don't answer controversial questions aswell."
GITA = "Give brief and concise explanation for the verses of Bhagvat Gita that I ask."
REFERENCE = "Provide references to these following topics about Hinduism from authentic hinduism related sites, with little to no explanation"

def get_reply(question, option):
    response = client.chat.completions.create(
        model="gpt-3.5-turbo-16k",
        messages=[
            {
                "role": "system",
                "content": option
            },
            {
                "role": "user",
                "content": question
            }
        ],
        temperature=1,
        max_tokens=200,
        top_p=1,
        frequency_penalty=0,
        presence_penalty=0
    )

    return response.choices[0].message.content


def main():
    f = open("query.txt")
    query = f.readline()
    if sys.argv[1] == "g":
        print(get_reply(query, GITA))
    elif sys.argv[1] == "n":
        print(get_reply(query, NORMAL))
    elif sys.argv[1] == "r":
        print(get_reply(query, REFERENCE))


if __name__ == "__main__":
    main()
