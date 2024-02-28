<script setup lang="ts">
import { ref, defineEmits, watch } from 'vue'

const props = defineProps<{
    msg: string,
    id: string,
    correct: boolean
}>()

const emits = defineEmits<{
    (e: 'clicked', id: string): void
}>()

const iscorrect = ref(false)
const wronganswer = ref(false)

function checkcorrect() {
    if (props.correct.valueOf() == true) {
        iscorrect.value = true;
    }
    else {
        wronganswer.value = true;
    }
    emits('clicked', props.id);
}

watch(() => props.msg, () => {
    iscorrect.value = false
    wronganswer.value = false;
});

</script>

<template>
    <div :id="id.toString()" class="select-none quizelement grid content-center text-xl h-30 w-100 p-4" v-on:click="checkcorrect()"
        :class="{ correct: iscorrect, notcorrect: wronganswer }">
        <div class="select-none">{{ msg }}</div>
    </div>
</template>

<style scoped>
.quizelement {
    border-radius: 0.5cm;
    background: rgb(91, 147, 140);
}

.quizelement:hover {
    background: rgb(78, 183, 160);
}

.correct {
    background: lightgreen !important;
}

.notcorrect {
    background: rgb(194, 16, 16) !important;
}

.select-none {
    -webkit-user-select: none;
    /* Chrome, Safari, Opera */
    -moz-user-select: none;
    /* Firefox */
    -ms-user-select: none;
    /* Internet Explorer/Edge */
    user-select: none;
}</style>
