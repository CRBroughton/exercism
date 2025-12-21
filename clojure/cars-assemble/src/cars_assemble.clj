(ns cars-assemble)

(defn production-rate
  "Returns the assembly line's production rate per hour,
   taking into account its success rate"
  [speed]
  (->> speed                        ; Start with the speed value
       (* 221)                      ; => (* 221 speed) - Multiply by base rate
       (* (cond                     ; => (* success-rate result) - Apply success rate:
            (= speed 0) 0.0         ;   - Speed 0: 0% success
            (<= 1 speed 4) 1.0      ;   - Speed 1-4: 100% success
            (<= 5 speed 8) 0.9      ;   - Speed 5-8: 90% success
            (= speed 9) 0.8         ;   - Speed 9: 80% success
            (= speed 10) 0.77       ;   - Speed 10: 77% success
            :else 0.0))))

(defn working-items
  "Calculates how many working cars are produced per minute"
  [speed]
  (-> speed                         ; Start with the speed value
      production-rate               ; => (production-rate speed) - Get hourly rate
      (/ 60)                        ; => (/ result 60) - Convert to per minute
      int))                         ; => (int result) - Convert to integer
