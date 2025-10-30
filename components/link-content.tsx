"use client"

import Image from "next/image"
import Link from "next/link"

import { HoverEffect } from "@/components/ui/card-hover-effect"
import { useConfigStore } from "@/stores"

export function LinkContent() {
  const { categories } = useConfigStore()

  return (
    <div className="w-full min-h-screen pb-32 pt-8">
      <div id="main" className="mx-auto w-full max-w-7xl px-6 md:px-8">
        {categories.map((category, index) => {
          return (
            <div id={String(index)} key={index} className="mb-16 scroll-mt-20">
              <div className="mb-6">
                <div className="flex items-center gap-4">
                  <h2 className="text-3xl font-bold text-gray-900 dark:text-gray-50">
                    {category.title}
                  </h2>
                  <div className="h-px flex-1 bg-gradient-to-r from-gray-300 via-gray-200 to-transparent dark:from-gray-700 dark:via-gray-800" />
                </div>
              </div>
              <HoverEffect
                items={category.items.map(({ title, desc, link, icon }) => ({
                  link,
                  title,
                  description: desc,
                  icon
                }))}
              />
            </div>
          )
        })}
      </div>
    </div>
  )
}
