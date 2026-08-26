<script lang="ts">
    const faqs = [
        {
            q: "is this free?",
            a: "Yep, totally free, same as every other Hack Club event.",
        },
        {
            q: "what can I ship?",
            a: "Pick something that runs over HTTP to build. A site, a tool, a game, whatever you're into. You get real rewards for actually finishing it.",
        },
        {
            q: "am I eligible?",
            a: "If you're 13 to 18, you're in.",
        },
        {
            q: "do I need Hackatime?",
            a: "Yes. It's how we know how long you spent, and it's what your rewards are worked out from.",
        },
    ];
</script>

<section id="questions" class="w-full bg-tea-background px-4 py-20 md:py-28">
    <div class="mx-auto max-w-3xl">
        <h2 class="reveal font-instrument text-4xl text-tea-dark-olive md:text-6xl">
            ✦ questions
        </h2>

        <div class="mt-10 flex flex-col gap-4 md:mt-14">
            {#each faqs as faq (faq.q)}
                <details
                    class="reveal group rounded-3xl border-2 border-tea-dark-olive bg-tea-pale-green/40 open:bg-tea-pale-green/60"
                >
                    <summary
                        class="flex cursor-pointer list-none items-center justify-between gap-4 px-6 py-5 font-instrument text-xl text-tea-dark-olive md:text-2xl [&::-webkit-details-marker]:hidden"
                    >
                        {faq.q}
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2.5"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="shrink-0 transition-transform duration-200 group-open:rotate-45"
                        >
                            <path d="M12 5v14M5 12h14" />
                        </svg>
                    </summary>
                    <p class="px-6 pb-5 text-lg text-tea-dark-olive">{faq.a}</p>
                </details>
            {/each}
        </div>
    </div>
</section>
