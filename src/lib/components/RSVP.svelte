<script lang="ts">
    import Plus from "@lucide/svelte/icons/plus";
</script>

<button
    class="outer group flex cursor-pointer items-center justify-center rounded-full border-none bg-tea-dark-olive"
>
    <span
        class="
        inner box-border inline-flex items-center gap-2
        text-3xl md:text-4xl text-tea-yellow-green font-instrument
        border-2 border-tea-dark-olive bg-tea-olive border-sol
        py-2.5 px-12 rounded-full
        ">RSVP Now <Plus /></span
    >
</button>

<style>
    .inner {
        transform: translateY(-0.2em);
        transition: transform 0.2s ease;
    }

    .outer:hover .inner {
        /* pull upwards */
        transform: translateY(-0.33em);
    }

    .outer:active .inner {
        /* push downwards */
        transform: translateY(0);
    }
</style>
