from openai import OpenAI
import sys

client = OpenAI()
NORMAL = "I, Arjuna will ask questions, and you, as Krishna, will briefly respond in alignment with the teachings of the Bhagavad Gita."
REFERENCE = "Give brief and concise explanation for the verses of Bhagvat Gita that I ask."

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
    if sys.argv[1] == "n":
        print(get_reply(query, NORMAL))
    elif sys.argv[1] == "r":
        print(get_reply(query, REFERENCE))


if __name__ == "__main__":
    main()
