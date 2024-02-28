<script setup lang="ts">
import { ref } from 'vue'
import type { Ref } from 'vue'
import QuizOption from './QuizOption.vue';
import QuizGrid41 from './QuizGrid41.vue'
import Question from './Question.vue';
import { QuestionHandler, QuestionItem } from '../questions';


defineProps<{ msg: string }>()



const Questions = new QuestionHandler()

let questiontext = ref("Wie heiße ich?");
const answers: Ref<QuestionItem[]> = ref([
  { text: "Maria", correct: false },
  { text: "Joche", correct: false },
  { text: "Liz", correct: false },
  { text: "Brandon", correct: true },
])

let currentwinner = 4;
let points = 0;

function checkforpoints(id: string) {
  if (parseInt(id) === currentwinner) {
    points++;
  }
  setTimeout(
      getNewAnswers,
      1000);
}

async function getNewAnswers() {
  let qs = await Questions.getNewQuestion();
  questiontext.value = qs.text
  answers.value = qs.answers

  for (let index = 0; index < qs.answers.length; index++) {
    if (qs.answers[index].correct) {
      currentwinner = index + 1
      break
    }
  }
}


</script>

<template>
  <div class="container-xl m-4">
    <div class="text-2xl text-center my-6 ">
      Points : <span class="text-blue-600">{{ points }}</span>
    </div>
    
   
    <Question :msg="questiontext" ></Question>
    <QuizGrid41>
      <template v-slot:one>
        <QuizOption id="1" :msg="answers[0].text" :correct="answers[0].correct" @clicked="checkforpoints">

        </QuizOption>
      </template>
      <template v-slot:two>
        <QuizOption id="2" :msg="answers[1].text" :correct="answers[1].correct" @clicked="checkforpoints">

        </QuizOption>
      </template>
      <template v-slot:three>
        <QuizOption id="3" :msg="answers[2].text" :correct="answers[2].correct" @clicked="checkforpoints">

        </QuizOption>
      </template>
      <template v-slot:four>
        <QuizOption id="4" :msg="answers[3].text" :correct="answers[3].correct" @clicked="checkforpoints">

        </QuizOption>
      </template>
    </QuizGrid41>

  </div>
</template>

<style scoped></style>
