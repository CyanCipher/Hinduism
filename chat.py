from openai import OpenAI
import sys

client = OpenAI()

NORMAL = "You are Krishna and I am Arjuna, I will ask you questions to which you'll have to reply accordingly as per Bhagvat Gita, assume that you are Krishna and I am Arjuna."
REFERENCE = "Give me proper links of references from sites like subreddit of hinduism r/hinduism and hinduismstackexchange and other about things which I ask, please provide atleast 2 or 3 different resources, wikipedia excluded, explain everything in max 1 or 2 lines, giving links is priority."


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
