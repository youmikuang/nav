"use client";

import { LinkContent } from "@/components/link-content"
import { Sidebar } from "@/components/sidebar"
import { SiteFooter } from "@/components/site-footer"
import { SiteHeader } from "@/components/site-header"
import { useState } from 'react';
import { cn } from "@/lib/utils"
import { AnimatePresence, motion } from "framer-motion"

export default function IndexPage() {
  const [isSidebarOpen, setSidebarOpen] = useState(false);

  return (
    <div className="container relative mx-auto min-h-screen w-full px-0">
      <AnimatePresence>
        {isSidebarOpen && (
          <motion.div
            onClick={() => setSidebarOpen(false)}
            className="fixed inset-0 z-30 bg-black/50 sm:hidden"
            aria-hidden="true"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.2 }}
          />
        )}
      </AnimatePresence>
      <div className="flex">
        <div
          className={cn(
            "fixed z-40 h-screen w-64 transition-transform duration-300 ease-in-out sm:translate-x-0",
            isSidebarOpen ? "translate-x-0" : "-translate-x-full"
          )}
        >
          <Sidebar onLinkClick={() => setSidebarOpen(false)} />
        </div>

        <div className="w-full sm:pl-64">
          <SiteHeader onMenuClick={() => setSidebarOpen(!isSidebarOpen)} />
          <LinkContent />
          <SiteFooter />
        </div>
      </div>
    </div>
  )
}
